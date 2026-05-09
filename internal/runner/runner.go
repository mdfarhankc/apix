package runner

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/mdfarhankc/apix/internal/client"
	"github.com/mdfarhankc/apix/internal/config"
	"github.com/mdfarhankc/apix/internal/formatter"
)

type Options struct {
	Method      string
	RawURL      string
	Body        []byte
	RawHeaders  []string
	RawQuery    []string
	ContentType string
}

func Run(opts Options) {
	if strings.HasPrefix(string(opts.Body), "@") {
		path := string(opts.Body)[1:]
		data, err := os.ReadFile(path)
		if err != nil {
			formatter.Fail(err)
		}
		opts.Body = data
	}

	resolved, err := config.Resolve(opts.RawURL)
	if err != nil {
		formatter.Fail(err)
	}

	if len(opts.RawQuery) > 0 {
		parsed, err := url.Parse(resolved.URL)
		if err != nil {
			formatter.Fail(err)
		}
		q := parsed.Query()
		for _, kv := range opts.RawQuery {
			k, v, _ := strings.Cut(kv, "=")
			q.Add(k, v)
		}
		parsed.RawQuery = q.Encode()
		resolved.URL = parsed.String()
	}

	headers := client.ParseHeaders(opts.RawHeaders)
	for k, v := range resolved.Headers {
		if _, ok := headers[k]; !ok {
			headers[k] = v
		}
	}

	if opts.ContentType != "" {
		if _, ok := headers["Content-Type"]; !ok {
			headers["Content-Type"] = opts.ContentType
		}
	}

	fmt.Printf(
		"%s %s\n\n",
		color.CyanString(opts.Method),
		resolved.URL,
	)

	resp, err := client.Do(client.Request{
		Method:  opts.Method,
		URL:     resolved.URL,
		Body:    opts.Body,
		Headers: headers,
	})
	if err != nil {
		formatter.Fail(err)
	}

	formatter.PrintResponse(resp)
}
