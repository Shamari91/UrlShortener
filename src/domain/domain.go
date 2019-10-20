package domain

type UrlDataRepository interface {
	Store(urlData UrlData)
	RetrieveLongUrlFrom(longUrl string)
}

type UrlData struct {
	shortUrl string
	longUrl  string
}
