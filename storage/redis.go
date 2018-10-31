package storage

import (
	"net/url"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
	"github.com/poccariswet/shorterql/hash"
)

var Pool *redis.Pool

func NewPool(redisAddr string) *redis.Pool {
	redisPool := redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", redisAddr)
		if err != nil {
			return nil, err
		}

		return c, err
	}, 10)

	return redisPool
}

func LoadAndCountUp(id string) (string, error) {
	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("HINCRBY", id, "count", 1)
	if err != nil {
		return "", errors.Wrapf(err, "HINCRBY %s count up err", id)
	}

	longURL, err := redis.String(conn.Do("HGET", id, "long_url"))
	if err != nil {
		return "", errors.Wrapf(err, "HGET %s longurl err", id)
	}

	return longURL, nil
}

func (u *UrlSho) Set() error {
	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("HSET", u.ID, "id", u.ID)
	if err != nil {
		return errors.Wrapf(err, "HSET %s long_url err", u.ID)
	}

	_, err = conn.Do("HSET", u.ID, "long_url", u.LongURL)
	if err != nil {
		return errors.Wrapf(err, "HSET %s long_url err", u.ID)
	}

	_, err = conn.Do("HSET", u.ID, "short_url", u.ShortURL)
	if err != nil {
		return errors.Wrapf(err, "HSET %s short_url err", u.ID)
	}

	_, err = conn.Do("HSET", u.ID, "count", u.Count)
	if err != nil {
		return errors.Wrapf(err, "HSET %s count err", u.ID)
	}

	_, err = conn.Do("HSET", u.ID, "createdAt", u.CreatedAt)
	if err != nil {
		return errors.Wrapf(err, "HSET %s created at err", u.ID)
	}

	return nil
}

func SaveURL(longURL string) (*UrlSho, error) {
	u := &UrlSho{
		ID:        hash.CreateHashID(),
		LongURL:   longURL,
		Count:     0,
		CreatedAt: time.Now().String(),
	}
	u.ShortURL = CreateURL(u.ID)
	if err := u.Set(); err != nil {
		return nil, errors.Wrap(err, "redis url set err")
	}

	return u, nil
}

func FetchURLInfoByID(id string) (*UrlSho, error) {
	conn := Pool.Get()
	defer conn.Close()

	var u UrlSho
	v, err := redis.Values(conn.Do("HGETALL", id))
	if err != nil {
		return nil, errors.Wrapf(err, "redis HGETALL %s err", id)
	}

	if err := redis.ScanStruct(v, &u); err != nil {
		return nil, errors.Wrap(err, "redis scan struct err")
	}

	return &u, nil
}

// TODO fetch all list
func FetchURLInfoList() ([]UrlSho, error) {
	return []UrlSho{}, nil
}

func CreateURL(id string) string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	u := &url.URL{
		Scheme: "http",
		Host:   "localhost:" + port,
		Path:   id,
	}

	return u.String()
}
