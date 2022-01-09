package app

import "github.com/miramariadev/go-musthave-shortener-tpl/internal/app/helpers"

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
	id := string(rune(helpers.GenerateInteger())) + helpers.GenerateString()
	srv.storage.AddURL(id, url)
	shortedURL := url + "/" + id
	return shortedURL
}

func (srv *URLShortenerService) GetLongURLById(id string) (string, error) {
	url, err := srv.storage.GetURL(id)
	if err != nil {
		return "", err
	}

	return url, nil
}
