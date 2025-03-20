package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
)

var client = http.Client{}

func callBoth(ctx context.Context, errVal string, slowURL string, fastURL string) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		err := callServer(ctx, "slow", slowURL)
		if err != nil {
			cancel()
		}
	}()

	go func() {
		defer wg.Done()

		err := callServer(ctx, "fast", fastURL+"?error="+errVal)
		if err != nil {
			cancel()
		}
	}()

	wg.Wait()
	fmt.Println("Complete both")
}

func callServer(ctx context.Context, label string, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Request error about ", label, ":", err)
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Response error about ", label, ":", err)
		return err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error about ", label, ":", err)
		return err
	}

	result := string(data)
	if result != "" {
		fmt.Println("Result of ", label, ":", result)
	}
	if result == "error" {
		fmt.Println("Cancel:", label)
		return errors.New("Error occurred")
	}
	return nil
}
