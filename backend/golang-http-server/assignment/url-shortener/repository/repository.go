package repository

import (
	"sync"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
)

type URLRepository struct {
	mu   sync.Mutex
	Data map[string]string
}

func NewMapRepository() URLRepository {
	data := make(map[string]string, 0)
	return URLRepository{
		Data: data,
	}
}

func (r *URLRepository) Get(path string) (*entity.URL, error) {
	value, ok := r.Data[path]
	if ok {
		return &entity.URL{
			ShortURL: path,
			LongURL:  value,
		}, nil
	}
	return &entity.URL{}, entity.ErrURLNotFound
}


func (r *URLRepository) Create(longURL string) (*entity.URL, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	path := entity.GetRandomShortURL(longURL)
	r.Data[path] = longURL
	return &entity.URL{LongURL: longURL, ShortURL: path}, nil
}

func (r *URLRepository) CreateCustom(longURL, customPath string) (*entity.URL, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	if _, ok := r.Data[customPath]; ok {
		return nil, entity.ErrCustomURLIsExists
	}
	r.Data[customPath] = longURL
	return &entity.URL{
		ShortURL: customPath,
		LongURL:  longURL,
	}, nil
}
