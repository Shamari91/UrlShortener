package interfaces

import (
	"domain"
)

type CacheHandler interface {
	Set(key, value string) error
	Get(key string) (string, error)
}

type CacheRepo struct {
	cacheHandler CacheHandler
}

type CacheUrlRepo CacheRepo

func NewCacheUrlRepo(cacheHandler CacheHandler) *CacheUrlRepo {
	cacheUrlRepo := new(CacheUrlRepo)
	cacheUrlRepo.cacheHandler = cacheHandler
	return cacheUrlRepo
}

func (repo *CacheUrlRepo) Store(urlData *domain.UrlData) {
	repo.cacheHandler.Set(urlData.LongUrl, urlData.ShortUrl)
}

func (repo *CacheUrlRepo) RetrieveLongUrlFrom(shortUrl string) (string, error) {
	return repo.cacheHandler.Get(shortUrl)
}
