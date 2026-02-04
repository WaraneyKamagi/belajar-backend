package handlers

import (
	"fmt"
	"net/http"
)

func Handlercoba(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Halo")
}
