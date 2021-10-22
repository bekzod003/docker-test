package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/docker", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello from docker")
	})

	http.ListenAndServe(":8080", nil)
}
