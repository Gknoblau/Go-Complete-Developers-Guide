package main

import (
	"io"
	"log"
	"os"
)

func main() {
	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, f)
}
