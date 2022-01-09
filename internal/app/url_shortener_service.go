package app

import (
	"errors"
	"math/rand"
	"strconv"
)

const maxInt = 99999

type Storage interface {
	AddURL(id string, url string)
	GetURL(id string) string
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
	return id
}

func (srv *URLShortenerService) GetLongURLByID(id string) (string, error) {
	url := srv.storage.GetURL(id)
	if url == "" {
		return url, errors.New("requested link does not exist")
	}

	return url, nil
}
