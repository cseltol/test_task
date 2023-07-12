package main

import (
	"net/http"
	"test_task/config"
)

func Start(cfg config.Config) error {
	srv := NewServer()
	return http.ListenAndServe(cfg.BindAddr, srv)
}