package main

import "net/http"

func router() {
	http.HandleFunc("/", cfg.db.home)
	http.HandleFunc("/login", cfg.db.login)
	http.HandleFunc("/register", cfg.db.register)
	http.HandleFunc("/dashboard", cfg.db.dashboard)
}
