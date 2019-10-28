package verify

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	tokenWithBearer    = `Bearer eyJhbGciOiJIUzI1NiJ9.gYp-4ieS6GdXCU3QOduonlmLAU7E1Tz-DdksGX0a6SinuXLOSqos2Q6zBJiSze7leabHwQQT6DE4qGWtX4dHyzr5CfZZ_HYvU9d7bYRAAIVlCZttEcR4B5q_2SN1KOjD7Ly_I5j6btG0VoXwcNqtZqfxXNUeL9kc4SVVx9g6tFI.CKFYwLEzx63vtmFcWJGIZa1nLllmlpZquIUfhZI7h4E`
	tokenWithOutBearer = `eyJhbGciOiJIUzI1NiJ9.gYp-4ieS6GdXCU3QOduonlmLAU7E1Tz-DdksGX0a6SinuXLOSqos2Q6zBJiSze7leabHwQQT6DE4qGWtX4dHyzr5CfZZ_HYvU9d7bYRAAIVlCZttEcR4B5q_2SN1KOjD7Ly_I5j6btG0VoXwcNqtZqfxXNUeL9kc4SVVx9g6tFI.CKFYwLEzx63vtmFcWJGIZa1nLllmlpZquIUfhZI7h4E`
)

func TestVerifyShouldBeSuccess(t *testing.T) {
	t.Run("verify should be success", func(t *testing.T) {
		expected := 200
		e := echo.New()
		handler := func(c echo.Context) error {
			return c.String(http.StatusOK, "test")
		}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, tokenWithBearer)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)

		h := LineJWT()(handler)

		if assert.NoError(t, h(c), "verify code error") {
			actual := c.Response().Status
			assert.Equal(t, expected, actual, "verify code should be %v but get &v", expected, actual)
		}
	})
}

func TestVerifyShouldNotExist(t *testing.T) {
	t.Run("verify should not exist", func(t *testing.T) {
		expected := "The access token not exist"
		e := echo.New()
		handler := func(c echo.Context) error {
			return c.String(http.StatusOK, "test")
		}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, tokenWithOutBearer)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)

		h := LineJWT()(handler)
		actual := h(c).(*echo.HTTPError).Message
		assert.Equal(t, expected, actual, "verify message should be : %v, but get : &v", expected, actual)
	})
}

func TestIsValidChannelID(t *testing.T) {
	t.Run("valid channel id should be true", func(t *testing.T) {
		expected := true
		actual := isValidChannelID(channelID)
		assert.Equal(t, expected, actual, "valid channel id should be : %v, but get : &v", expected, actual)
	})

	t.Run("valid channel id should be false", func(t *testing.T) {
		expected := false
		actual := isValidChannelID("fakeChannelID")
		assert.Equal(t, expected, actual, "valid channel id should be : %v, and : &v", expected, actual)
	})
}

func TestGetTokenShouldBeSuccess(t *testing.T) {
	t.Run("get token should be success", func(t *testing.T) {
		token := tokenWithBearer
		expected := tokenWithOutBearer
		actual := getToken(token)
		if actual != expected {
			t.Errorf("get token should be %v but get %v", expected, actual)
		}
	})

	t.Run("get token should be empty", func(t *testing.T) {
		token := ``
		expected := ""
		actual := getToken(token)
		if actual != expected {
			t.Errorf("get token should be %v but get %v", expected, actual)
		}
	})
}
