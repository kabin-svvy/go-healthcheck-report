package verify

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// Handler .
type Handler struct {
	config *viper.Viper
}

// ResponseVerify .
type ResponseVerify struct {
	ClientID         string `json:"client_id,omitempty"`
	ExpiresIn        int    `json:"expires_in,omitempty"`
	Scope            string `json:"scope,omitempty"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

// Response .
type Response struct {
	Error string `json:"error,omitempty"`
}

var (
	verifyURI     = "https://api.line.me/oauth2/v2.1/verify?access_token=%s"
	authorization = "Authorization"
	channelID     = "1653377896"
)

// LineJWT .
func LineJWT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := getToken(c.Request().Header.Get(authorization))
			resVer := ResponseVerify{}
			url := fmt.Sprintf(verifyURI, token)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				return &echo.HTTPError{
					Code:     http.StatusInternalServerError,
					Message:  "failed to get request",
					Internal: err,
				}
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				return &echo.HTTPError{
					Code:     http.StatusInternalServerError,
					Message:  "failed to send request",
					Internal: err,
				}
			}

			if res != nil {
				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					return &echo.HTTPError{
						Code:     http.StatusInternalServerError,
						Message:  "failed to read request body",
						Internal: err,
					}
				}

				err = json.Unmarshal(body, &resVer)
				if err != nil {
					return &echo.HTTPError{
						Code:     http.StatusInternalServerError,
						Message:  "failed to unmarshal request body",
						Internal: err,
					}
				}

				if res.StatusCode != http.StatusOK {
					err := errors.New(resVer.Error)
					return &echo.HTTPError{
						Code:     res.StatusCode,
						Message:  resVer.ErrorDescription,
						Internal: err,
					}
				}

				if !isValidChannelID(resVer.ClientID) {
					err := errors.New("invalid jwt")
					return &echo.HTTPError{
						Code:     http.StatusUnauthorized,
						Message:  err.Error(),
						Internal: err,
					}
				}

				if isExpired(resVer.ExpiresIn) {
					err := errors.New("expired jwt")
					return &echo.HTTPError{
						Code:     http.StatusUnauthorized,
						Message:  err.Error(),
						Internal: err,
					}
				}
			}
			return next(c)
		}
	}
}

func isExpired(expiresIn int) bool {
	if expiresIn > 0 {
		return false
	}
	return true
}

func isValidChannelID(clientID string) bool {
	if strings.Compare(clientID, channelID) != 0 {
		return false
	}
	return true
}

func getToken(authVal string) string {
	s := strings.Split(authVal, " ")
	if len(s) > 1 {
		return s[1]
	}
	return ""
}
