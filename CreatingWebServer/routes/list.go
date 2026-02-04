package routes

import (
	"creatingwebserver/handlers"
	"net/http"
)

func RegisterRouters() {
	http.HandleFunc("halo", handlers.Handlercoba)
	http.HandleFunc("/", handlers.HandlerUtama)
	http.HandleFunc("/login", handlers.HandlerLogin)
	http.ListenAndServe(":8080", nil)
}
