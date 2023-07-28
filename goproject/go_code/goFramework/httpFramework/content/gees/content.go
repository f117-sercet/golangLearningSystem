package gees

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Content struct {

	// origin objets
	Writer http.ResponseWriter
	Req    *http.Request
	// request info

	Path   string
	Method string

	// response info

	StatusCode int
}

func newContent(w http.ResponseWriter, req *http.Request) *Content {

	return &Content{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}
func (c *Content) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Content) Query(key string) string {

	return c.Req.URL.Query().Get(key)
}

func (c *Content) Status(code int) {

	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Content) SetHeader(key string, value string) {

	c.Writer.Header().Set(key, value)
}
func (c *Content) String(code int, format string, values ...interface{}) {

	c.SetHeader("content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c Content) JSON(code int, obj interface{}) {

	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)

	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}
func (c *Content) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}
func (c *Content) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
