# Go checkitout [![Go](https://github.com/Leoff00/gocheckitout/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/Leoff00/gocheckitout/actions/workflows/build.yml)

Go checkitout is project to make requests and bring site/webpage infos like:

```Go
type Custom struct {
	Header     http.Header // Header contents like content-type, accept, etc...
	StatusCode int16 //output default status code - 200. (This app only returns 200 if OK, if don't will return the specified configurated status code from errors.)
	RawBody    *io.ReadCloser //The raw body not converted in slice of byte, giving an options to optimize the request body as you want
	Body       *[]byte // The body converted in slice of bytes.
	Url        *url.URL //the own url of the request that you are specifying in the dummy text file.
	Timestamp  string //Default timestamp (IN BRAZILIAN TIMEZONE) of the moment of request.
}

```
---
**To run this this project type in the terminal:**

## if you want to compile:

## see the [build.sh](https://github.com/Leoff00/gocheckitout/blob/dev_1.1.0/build.sh) file

**or**

```sh
go run main.go
```

**and if you want to make benchmark of the project**

```sh
time go run main.go
```

### NOTE: this is ONE personal average benchmark of the changelog.

# 

# CHANGELOG

[See here](https://github.com/Leoff00/gocheckitout/releases) the changelog of the all tags/releases.

#

# Others

This is a side project to improve my skills in go lang. Learn Golang here in the [Golang Docs](https://go.dev/doc/).

