package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/shorten", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Shorten")
	})

	http.HandleFunc("/expand", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Expand")
	})
	http.ListenAndServe(":8080", nil)
}
