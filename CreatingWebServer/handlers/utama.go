package handlers

import (
	"fmt"
	"net/http"
)

func HandlerUtama(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
