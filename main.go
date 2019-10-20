package main

import (
	"interfaces"
	"net/http"
	"usecases"
)

func main() {
	urlDataInteractor := new(usecases.UrlDataInteractor)
	webserviceHandler := interfaces.WebServiceHandler{}
	webserviceHandler.UrlDataInteractor = urlDataInteractor

	http.HandleFunc("/shorten", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.CreateShortUrl(res, req)
	})

	http.HandleFunc("/expand", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.ExpandShortUrl(res, req)
	})
	http.ListenAndServe(":8080", nil)
}
