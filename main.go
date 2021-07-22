package main

import (
	"log"

	"github.com/evelritual/goose"
)

func main() {
	err := goose.Run(nil)
	if err != nil {
		log.Fatal(err)
	}
}
