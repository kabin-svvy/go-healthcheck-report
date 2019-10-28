package report

import (
	"log"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type Context struct {
	echo.Context
	code     int
	response interface{}
}

func (Context) Bind(i interface{}) error {
	*i.(*Request) = Request{
		TotalWebsites: 10,
		Success:       7,
		Failure:       3,
		TotalTime:     100000,
	}
	return nil
}

func (c *Context) JSON(code int, i interface{}) error {
	c.code = code
	c.response = i
	return nil
}

func TestCreateShouldBeSuccess(t *testing.T) {
	t.Run("test create should be success", func(t *testing.T) {
		expected := 200
		c := &Context{}
		err := Create(c)
		if err != nil {
			log.Println(err)
		}
		actual := c.code
		// res := c.response.(Response)
		// actual := res.Message
		assert.Equal(t, expected, actual, "test create should be %v but get %v", expected, actual)
	})
}

func TestValidateRequest(t *testing.T) {
	t.Run("total websites should equal or more than zero", func(t *testing.T) {
		req := Request{
			TotalWebsites: 0,
		}
		actual := req.validate()
		if actual != nil {
			t.Errorf("total websites should equal or more than zero but get error")
		}
	})

	t.Run("total websites should get error", func(t *testing.T) {
		req := Request{
			TotalWebsites: -1,
		}
		actual := req.validate()
		if actual == nil {
			t.Errorf("total websites should get error but get nil")
		}
	})

	t.Run("success should equal or more than zero", func(t *testing.T) {
		req := Request{
			Success: 0,
		}
		actual := req.validate()
		if actual != nil {
			t.Errorf("success should equal or more than zero but get error")
		}
	})

	t.Run("success should get error", func(t *testing.T) {
		req := Request{
			Success: -1,
		}
		actual := req.validate()
		if actual == nil {
			t.Errorf("success should get error but get nil")
		}
	})

	t.Run("failure should equal or more than zero", func(t *testing.T) {
		req := Request{
			Failure: 0,
		}
		actual := req.validate()
		if actual != nil {
			t.Errorf("failure should equal or more than zero but get error")
		}
	})

	t.Run("failure should get error", func(t *testing.T) {
		req := Request{
			Failure: -1,
		}
		actual := req.validate()
		if actual == nil {
			t.Errorf("failure should get error but get nil")
		}
	})

	t.Run("success and failure should equal or less than total websites", func(t *testing.T) {
		req := Request{
			TotalWebsites: 10,
			Success:       7,
			Failure:       3,
		}
		actual := req.validate()
		if actual != nil {
			t.Errorf("success and failure should equal or more than total websites but get error")
		}
	})

	t.Run("success and failure should get error", func(t *testing.T) {
		req := Request{
			TotalWebsites: 10,
			Success:       7,
			Failure:       4,
		}
		actual := req.validate()
		if actual == nil {
			t.Errorf("success and failure should get error but get nil")
		}
	})

	t.Run("total time should more than zero", func(t *testing.T) {
		req := Request{
			TotalTime: 10000,
		}
		actual := req.validate()
		if actual != nil {
			t.Errorf("success should more than zero but get error")
		}
	})

	t.Run("total time should get error", func(t *testing.T) {
		req := Request{
			TotalTime: -1,
		}
		actual := req.validate()
		if actual == nil {
			t.Errorf("total time should get error but get nil")
		}
	})
}
