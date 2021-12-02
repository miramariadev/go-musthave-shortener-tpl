package app

import (
	"math/rand"
	"strconv"
)

type Storage interface {
	AddURL(id string, url string)
	GetURL(id string) (string, error)
}

type URLShortenerService struct {
	storage Storage
}

func NewURLShortenerService(storage Storage) *URLShortenerService {
	return &URLShortenerService{
		storage: storage,
	}
}

func (srv *URLShortenerService) CreateShortURL(url string) string {
	id := strconv.Itoa(rand.Intn(maxInt))
	srv.storage.AddURL(id, url)
	shortedURL := "http://localhost:8080/" + id //url + "/" + id
	return shortedURL
}

func (srv *URLShortenerService) GetLongURLByID(id string) (string, error) {
	url, err := srv.storage.GetURL(id)
	if err != nil {
		return "", err
	}

	return url, nil
}
