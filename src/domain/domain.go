package domain

type UrlDataRepository interface {
	Store(urlData UrlData)
	FindByLongUrl(longUrl string)
}

type UrlData struct {
	shortUrl string
	longUrl  string
}
