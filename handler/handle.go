package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/poccariswet/shorterql/storage"
)

type URL struct {
	Url string `json:"url" validate:"required,url"`
}

func RedirectHandler(c echo.Context) error {
	code := c.Param("id")
	url, err := storage.LoadAndCountUp(code)
	if err != nil {
		return c.JSON(http.StatusBadRequest, CustomBadResponse(err))
	}

	return c.Redirect(http.StatusMovedPermanently, url)
}

func UrlShorterStatusHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "UrlShorterStatusHandler")
}

func UrlShorterHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "UrlShorterHandler")
}
