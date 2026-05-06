package client

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

type Response struct {
	StatusCode int
	Status     string
	Duration   time.Duration
	Body       []byte
	Size       int
}

func Do(req Request) (*Response, error) {
	start := time.Now()

	httpReq, err := http.NewRequest(
		req.Method,
		req.URL,
		bytes.NewBuffer(req.Body),
	)
	if err != nil {
		return nil, err
	}

	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}

	client := &http.Client{}

	resp, err := client.Do(httpReq)
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
		Status:     resp.Status,
		Duration:   time.Since(start),
		Body:       body,
		Size:       len(body),
	}, nil
}
