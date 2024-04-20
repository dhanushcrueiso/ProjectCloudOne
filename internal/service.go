package internal

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	base        = 62
)

var URLMap = make(map[string]string)
var basepath = "http://shortn.ly/"

func ShorternUrl(req RequestData) RequestData {
	originalURL := req.URL
	if _, exists := URLMap[originalURL]; !exists {
		// Generate shortened URL
		fmt.Println("check 2")
		shortenedURL := generateShortURL()
		URLMap[originalURL] = basepath + shortenedURL
	}
	fmt.Println("check 3")
	res := RequestData{
		URL: URLMap[originalURL],
	}
	return res
}

// generateShortURL generates a random shortened URL
func generateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	var shortURL string
	for i := 0; i < 7; i++ {
		shortURL += string(base62Chars[rand.Intn(base)])
	}
	return shortURL
}

func getKeyByValue(req RequestData) string {
	for k, v := range URLMap {
		if v == req.URL {
			return k
		}
	}
	return "" // Return an empty string if the value is not found
}
