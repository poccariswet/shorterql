package storage

import "time"

type UrlSho struct {
	ID        string    `json:"id" redis:"id"`
	LongURL   string    `json:"long_url" redis:"long_url"`
	Count     int       `json:"count" redis:"count"`
	CreatedAt time.Time `json:"createdAt" redis:"createdAt"`
}
