package server

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from Mongoose\n")
}
