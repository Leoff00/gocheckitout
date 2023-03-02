package helpers

import (
	"io/ioutil"
	"log"
	"net/http"

	utils "github.com/leoff00/gocheckitout/utils"
)

//Helper to make request with options and custom return.
func Requester(url string) (*utils.Custom, error) {
	res, err := http.Get(url)

	if err != nil {
		log.Panicf("Something has go wrong with %s. Error: %s", res.Request.URL, err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Panicf("Cannot Read the body of %s. Error: %s", res.Request.URL, err)
	}

	CustomResponse := utils.NewCustomResponse(res.Header, int64(res.StatusCode), &res.Body, &body, res.Request.URL)
	return CustomResponse, nil

}
