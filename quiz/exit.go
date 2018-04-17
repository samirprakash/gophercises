package main

import (
	"log"
	"os"
)

func exit(m string) {
	log.Println(m)
	os.Exit(1)
}
