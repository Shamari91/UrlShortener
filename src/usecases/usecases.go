package usecases

import (
	"domain"
)

type UrlDataInteractor struct {
	urlDataRepository domain.UrlDataRepository
}

func Shorten(url string) (string, error) {
	return string("sdfd"), nil
}

func Expand(url string) (string, error) {
	return string("sdfdf"), nil
}
