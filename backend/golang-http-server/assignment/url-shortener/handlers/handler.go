package handlers

import (
	// "errors"
	"net/http"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/repository"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	repo *repository.URLRepository
}

func NewURLHandler(repo *repository.URLRepository) URLHandler {
	return URLHandler{
		repo: repo,
	}
}

func (h *URLHandler) Get(c *gin.Context) {
	// TODO: answer here
	shortURL := c.Param("shorturl")
	url, err := h.repo.Get(shortURL)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.Writer.Header().Set("Location", url.LongURL)

	c.JSON(http.StatusFound, gin.H{
		"message": "Redirecting...",
	})
}

func (h *URLHandler) Create(c *gin.Context) {
	// TODO: answer here
	var url entity.URL
	err := c.ShouldBindJSON(&url)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ErrBadRequest)
	}

	entityURLresp, err := h.repo.Create(url.LongURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"Data": entityURLresp,
	})
}

func (h *URLHandler) CreateCustom(c *gin.Context) {
	// TODO: answer here
	var url entity.URL
	err := c.ShouldBindJSON(&url)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ErrBadRequest)
	}

	entityURLresp, err := h.repo.CreateCustom(url.LongURL, url.ShortURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"Data": entityURLresp,
	})
}
