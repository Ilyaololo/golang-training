package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type LogWriter struct{}

func (LogWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))

	return len(p), nil
}

func init() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	lw := LogWriter{}

	io.Copy(lw, resp.Body)
}
