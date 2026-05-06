package client

import (
	"io"
	"net/http"
	"time"
)

type Response struct {
	StatusCode int
	Duration   time.Duration
	Body       []byte
}

func Get(url string) (*Response, error) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode: resp.StatusCode,
		Duration:   time.Since(start),
		Body:       body,
	}, nil
}
