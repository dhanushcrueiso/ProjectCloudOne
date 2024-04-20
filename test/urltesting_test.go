package test

import (
	"fmt"
	"infracloud/internal"
	"testing"
)

func TestGetShortenURL(t *testing.T) {

	// Define a mock request data
	requestData := internal.RequestData{
		URL: "https://www.example.com",
	}

	res := internal.ShorternUrl(requestData)

	res2 := internal.ShorternUrl(requestData)

	if res == res2 {
		fmt.Println("successfull")
	}

}
