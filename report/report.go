package report

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Request .
type Request struct {
	TotalWebsites int   `json:"total_websites"`
	Success       int   `json:"success"`
	Failure       int   `json:"failure"`
	TotalTime     int64 `json:"total_time"`
}

// Response .
type Response struct {
	Message string `json:"message"`
}

// Create .
func Create(c echo.Context) error {
	req := &Request{}
	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Message: err.Error(),
		})
	}
	if err := req.validate(); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, Response{
		Message: "success",
	})
}

func (r Request) validate() error {
	if r.TotalWebsites < 0 {
		return errors.New("request total websites should equal or more than zero")
	}
	if r.Success < 0 {
		return errors.New("request success should equal or more than zero")
	}
	if r.Success > r.TotalWebsites {
		return errors.New("request success should equal or less than total websites")
	}
	if r.Failure < 0 {
		return errors.New("request failure should equal or more than zero")
	}
	if r.Failure > r.TotalWebsites {
		return errors.New("request failure should equal or less than total websites")
	}
	if r.Success+r.Failure > r.TotalWebsites {
		return errors.New("request success and failure should equal or less than total websites")
	}
	if r.TotalTime < 0 {
		return errors.New("request total time should more than zero")
	}
	return nil
}
