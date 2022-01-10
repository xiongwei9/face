package util

import (
	"fmt"
	"os"
)

func Exit(code int, msg string) {
	if msg != "" {
		fmt.Printf("[face] %s\n", msg)
	}
	os.Exit(code)
}
