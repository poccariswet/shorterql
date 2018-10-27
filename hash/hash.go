package hash

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

// unique id
func CreateHashID() string {
	id := uuid.New()
	hash := id.String()
	h := hash[len(hash)-1]
	hashID := hash[:8]

	if 48 <= h && h <= 52 {
		hashID = strings.ToUpper(hashID)
	} else if 53 <= h && h <= 57 {
		b := []byte(hashID)
		b[h-50] = h + 33
		hashID = string(b)
	} else if 97 <= h && h <= 107 {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(8)
		hashID = Replace(h-32, hashID, n)
	} else if 108 <= h && h <= 117 {
		hashID = hashID[:7]
	}

	return hashID
}

func Replace(to byte, hash string, n int) string {
	s := ""
	for i, v := range hash {
		if i == n {
			v = rune(to)
		}
		s += string(v)
	}
	return s
}
