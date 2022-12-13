package helpers

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	utils "github.com/leoff00/gocheckitout/utils"
)

//Helper to make request with options and custom return.
func Requester(url string) (*utils.Custom, error) {
	res, err := http.Get(url)

	currTime := time.Now().UTC().Local()
	body, _ := ioutil.ReadAll(res.Body)

	custom := utils.Custom{
		Header:     res.Header,
		StatusCode: int16(res.StatusCode),
		RawBody:    &res.Body,
		Body:       &body,
		Url:        res.Request.URL,
		Timestamp:  currTime.String(),
	}

	if err != nil {

		log.Default().Fatal("Failed to make request...")
		return &custom, err
	}

	return &custom, nil

}
