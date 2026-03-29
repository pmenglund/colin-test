package main

import (
	"fmt"
	"io"
)

func run(w io.Writer) error {
	_, err := fmt.Fprintln(w, statusMessage())
	return err
}

func statusMessage() string {
	return "colin-test is ready"
}
