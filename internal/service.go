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

func getKeyByValue(m map[string]string, value string) string {
	for k, v := range m {
		if v == value {
			return k
		}
	}
	return "" // Return an empty string if the value is not found
}
