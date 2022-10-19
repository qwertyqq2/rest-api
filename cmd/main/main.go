package main

import (
	"log"
	"test_go/internal/app/apiserver"
)

func main() {
	config := apiserver.GetConfig()

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
