package util

import (
	"os"
	"fmt"
	"time"
)

func StartupFatalErr(err error) {
	if err != nil {
		fmt.Printf("startup error: %s\n", err.Error())
		os.Exit(EXIT_STATUS_STARTUP_ERR)
	}
}

// milliseconds since January 1, 1970 UTC
func UnixMillSeconds() int64 {
	return time.Now().UnixNano() / 1e6
}

func SafeEmptyFunc(f func()) func() {
	return func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Fprintf(os.Stderr, "panic: %+v\n", r)
			}
		}()
		f()
	}
}
