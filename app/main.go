package main

import (
	"log"
	"test_task/config"
)


func main() {
	config := config.NewConfig()

	if err := Start(config); err != nil {
		log.Fatal(err)
	}
}
