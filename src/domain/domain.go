package domain

type UrlDataCache interface {
	Store(urlData *UrlData)
	RetrieveLongUrlFrom(shortUrl string) (string, error)
}

type UrlDataRepository interface {
	Store(urlData UrlData)
	RetrieveLongUrlFrom(shortUrl string)
}

type UrlData struct {
	ShortUrl string
	LongUrl  string
}
