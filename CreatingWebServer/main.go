package main

import (
	"creatingwebserver/routes"
	"net/http"
)

func main() {
	routes.RegisterRouters()
	http.ListenAndServe(":8080", nil)
}
