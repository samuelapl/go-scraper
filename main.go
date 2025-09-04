package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func fetchURL(url string) (string, error) {
	resp, err := http.Get(url) 
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL %s: %w", url, err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(bodyBytes), nil
}

func main() {
	urls := []string{
		"https://go.dev",
		"https://google.com",
		"https://wikipedia.org",
		"https://example.com",
		"https://golang.org/doc",
	}

	var wg sync.WaitGroup
	results := make(chan string, len(urls))

	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()
			content, err := fetchURL(u)
			if err != nil {
				results <- fmt.Sprintf("Error fetching %s: %v", u, err)
				return
			}
			results <- fmt.Sprintf("Fetched %s, length: %d", u, len(content))
		}(url)
	}

	wg.Wait()
	close(results)

	for result := range results { /
		fmt.Println(result)
	}
}