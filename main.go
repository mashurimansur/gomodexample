package main

import (
	"github.com/gosimple/slug"
	"log"
)

func main() {
	text := slug.Make("Ini Adalah Text")
	log.Println(text)
}