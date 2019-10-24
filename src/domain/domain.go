package domain

type UrlDataCache interface {
	Store(urlData UrlData)
	RetrieveLongUrlFrom(shortUrl string)
}

type UrlDataRepository interface {
	Store(urlData UrlData)
	RetrieveLongUrlFrom(shortUrl string)
}

type UrlData struct {
	LaongUrl string
	ShortUrl string
}
