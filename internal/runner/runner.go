package runner

import (
	"fmt"

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
	ContentType string
}

func Run(opts Options) {
	resolved, err := config.Resolve(opts.RawURL)
	if err != nil {
		formatter.Fail(err)
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
