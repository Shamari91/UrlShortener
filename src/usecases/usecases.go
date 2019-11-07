package usecases

import (
	"crypto/md5"
	"domain"
	"encoding/base64"
	"fmt"
	"regexp"
)

type UrlDataInteractor struct {
	urlDataCache      domain.UrlDataCache
	urlDataRepository domain.UrlDataRepository
}

func (interactor *UrlDataInteractor) Shorten(url string) string {
	shortUrl := CreateShortUrl(url)
	_, err := interactor.urlDataCache.RetrieveLongUrlFrom(shortUrl)
	if err != nil {
		interactor.urlDataCache.Store(&domain.UrlData{
			ShortUrl: shortUrl,
			LongUrl:  url,
		})
	}
	return shortUrl
}

func (interactor *UrlDataInteractor) Expand(shortUrl string) string {
	val, err := interactor.urlDataCache.RetrieveLongUrlFrom(shortUrl)
	if err != nil {
		return ""
	}
	return val
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
