package utils

import (
	"io"
	"net/http"
	"net/url"
)

// Customize the response with the way that you want
// But remember that respose must be legit in according within http struct.
type Custom struct {
	Header     http.Header
	StatusCode int16
	RawBody    *io.ReadCloser
	Body       *[]byte
	Url        *url.URL
	Timestamp  string
}
