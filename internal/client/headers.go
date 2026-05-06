package client

import "strings"

func ParseHeaders(headers []string) map[string]string {
	result := make(map[string]string)

	for _, header := range headers {
		parts := strings.SplitN(header, ":", 2)

		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		result[key] = value
	}

	return result
}
