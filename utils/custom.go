package utils

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

// Customize the response with the way that you want
// But remember that response must be legit in according within http struct.
type Custom struct {
	Header     http.Header    // Header contents like content-type, accept, etc...
	StatusCode int64          //output default status code - 200. (This app only returns 200 if OK, if don't will return the specified configurated status code from errors.)
	RawBody    *io.ReadCloser //The raw body not converted in slice of byte, giving an options to optimize the request body as you want
	Body       *[]byte        // The body converted in slice of bytes.
	Url        *url.URL       //the own url of the request that you are specifying in the dummy text file.
	Timestamp  time.Time      //Default timestamp (IN BRAZILIAN TIMEZONE) of the moment of request.
}

func NewCustomResponse(httpHeader http.Header, statusCode int64, rawBody *io.ReadCloser, body *[]byte, url *url.URL) *Custom {
	return &Custom{
		Header:     httpHeader,
		StatusCode: statusCode,
		RawBody:    rawBody,
		Body:       body,
		Url:        url,
		Timestamp:  time.Now().UTC().Local(),
	}
}
