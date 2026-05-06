package formatter

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/mdfarhankc/apix/internal/client"
)

func PrintResponse(resp *client.Response) {
	statusText := resp.Status

	switch {
	case resp.StatusCode >= 200 && resp.StatusCode < 300:
		statusText = color.GreenString(resp.Status)

	case resp.StatusCode >= 400 && resp.StatusCode < 500:
		statusText = color.YellowString(resp.Status)

	case resp.StatusCode >= 500:
		statusText = color.RedString(resp.Status)
	}

	fmt.Printf(
		"%s • %v • %d bytes\n\n",
		statusText,
		resp.Duration,
		resp.Size,
	)

	pretty, err := PrettyJSON(resp.Body)
	if err != nil {
		fmt.Println(string(resp.Body))
		return
	}

	fmt.Println(string(pretty))
}
