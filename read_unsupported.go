//go:build !windows

package main

import (
	"fmt"
	"runtime"
)

func read(text string) error {
	return fmt.Errorf("%s (%s architecture): %w", runtime.GOOS, runtime.GOARCH, errReadUnsupported)
}
