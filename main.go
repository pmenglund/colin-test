package main

import (
	"log"
	"os"
)

func main() {
	if err := run(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
