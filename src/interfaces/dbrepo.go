package interfaces

import (
	"domain"
	"fmt"
)

type DbHandler interface {
	Execute(statement string)
	Query(statement string) (Row, error)
}

type Row interface {
	Next() bool
	Scan(dest ...interface{})
}

type DbRepo struct {
	dbHandler DbHandler
}

type UrlDbRepo DbRepo

func NewUrlDbRepo(dbHandler DbHandler) *UrlDbRepo {
	urlRepo := new(UrlDbRepo)
	urlRepo.dbHandler = dbHandler
	return urlRepo
}

func (urlDbRepo *UrlDbRepo) Store(urlData *domain.UrlData) {
	urlDbRepo.dbHandler.Execute(fmt.Sprintf("INSERT INTO url (short_url, long_url) VALUES ('%v', '%v')", urlData.ShortUrl, urlData.LongUrl))
}

func (urlDbRepo *UrlDbRepo) FindByShortUrl(shortUrl string) {
	row, err := urlDbRepo.dbHandler.Query(fmt.Sprintf("SELECT long_url FROM url WHERE short_url = '%v' LIMIT 1", shortUrl))
	if err != nil {

	}

	var longUrl string
	row.Next()
	row.Scan(&longUrl)
}
