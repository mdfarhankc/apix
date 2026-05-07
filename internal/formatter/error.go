package formatter

import (
	"fmt"
	"os"
)

func Fail(err error) {
	fmt.Fprintln(os.Stderr, "Error:", err)
	os.Exit(1)
}
