package usecases

import (
	"crypto/md5"
	"domain"
	"encoding/base64"
	"fmt"
	"regexp"
)

type UrlDataInteractor struct {
	urlDataRepository domain.UrlDataRepository
}

func (interactor *UrlDataInteractor) Shorten(url string) string {
	/*shortUrl, err := interactor.urlDataRepository.RetrieveShortUrlFrom(url)
	if len(shortUrl) > 0 {
		return shortUrl
	}*/

	shortUrl := CreateShortUrl(url)
	//interactor.urlDataRepository.Store(UrlData{shortUrl, url})
	return shortUrl
}

func (interactor *UrlDataInteractor) RetrieveLongUrlFrom(shortUrl string) string {
	//return string("sfdf")
}

const ShortUrlPrefix = "ShortUrl/"

func CreateShortUrl(url string) string {
	hashedUrlStr := fmt.Sprintf("%x", md5.Sum([]byte(url)))
	urlWithoutSpecialCharacters, _ := RemoveSpecialCharactersFrom(hashedUrlStr)

	return ShortUrlPrefix + urlWithoutSpecialCharacters[:5]
}

func RemoveSpecialCharactersFrom(hashedUrl string) (string, error) {
	shortenedUrlRegex, err := regexp.Compile("[/+]")
	if err != nil {
		return "", err
	}

	shortenedUrl := shortenedUrlRegex.ReplaceAllString(base64.URLEncoding.EncodeToString([]byte(hashedUrl)), "_")
	return shortenedUrl, nil
}
