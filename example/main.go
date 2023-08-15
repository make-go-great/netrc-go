package main

import (
	"log"

	"github.com/make-go-great/netrc-go"
)

func main() {
	data, err := netrc.ParseFile("~/.netrc")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", data)
}
