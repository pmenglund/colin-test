package main

import (
	"bytes"
	"testing"
)

func TestRunWritesStatusMessage(t *testing.T) {
	var buf bytes.Buffer

	if err := run(&buf); err != nil {
		t.Fatalf("run() error = %v", err)
	}

	got := buf.String()
	want := "colin-test is ready\n"
	if got != want {
		t.Fatalf("run() output = %q, want %q", got, want)
	}
}
