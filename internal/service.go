package internal

import (
	"math/rand"
	"time"
)

const (
	base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	base        = 62
)

// generateShortURL generates a random shortened URL
func generateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	var shortURL string
	for i := 0; i < 7; i++ {
		shortURL += string(base62Chars[rand.Intn(base)])
	}
	return shortURL
}
