package helpers

import (
	"log"
	"net/http"
	"time"

	utils "github.com/leoff00/gocheckitout/utils"
)

//Helper to make request with options.
func Requester(url string) (*utils.Custom, error) {

	res, err := http.Get(url)

	custom := utils.Custom{
		Header:     res.Header,
		StatusCode: int16(res.StatusCode),
		Body:       &res.Body,
		Url:        res.Request.URL,
		Timestamp:  time.Now(),
	}

	if err != nil {
		log.Default().Fatal("Failed to make request...")
		return nil, err
	}

	return &custom, nil
}
