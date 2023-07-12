package main

import (
	"log"
	"net/http"
	"test_task/config"
)

func Start(cfg config.Config) error {
	srv := NewServer()
	err := http.ListenAndServe(cfg.BindAddr, srv)

	if err != nil {
		log.Println("FAILED TO START APP...")
		return err
	}
	log.Println("STARTED APP...")
	return nil
}