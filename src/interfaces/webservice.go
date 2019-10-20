package interfaces

import (
	"io"
	"net/http"
)

type UrlDataInteractor interface {
	Shorten(url string) string
	RetrieveLongUrlFrom(shortUrl string) string
}

type WebServiceHandler struct {
	UrlDataInteractor UrlDataInteractor
}

func (webServiceHandler WebServiceHandler) CreateShortUrl(res http.ResponseWriter, req *http.Request) {
	longUrl := req.FormValue("Url")
	if len(longUrl) == 0 {
		http.Error(res, "Url to shorten is missing", http.StatusBadRequest)
		return
	}

	shortUrl := webServiceHandler.UrlDataInteractor.Shorten(longUrl)
	io.WriteString(res, shortUrl)
}

func (webServiceHandler WebServiceHandler) ExpandShortUrl(res http.ResponseWriter, req *http.Request) {

}
