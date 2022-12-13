package api

import (
	"fmt"
	"log"
	"os"

	"github.com/leoff00/gocheckitout/helpers"
)

//Make request returning only the url and the status code.
func MakeRequest(url string) error {
	res, err := helpers.Requester(url)

	if err != nil {
		return err
	}

	if res.StatusCode == 200 {
		fmt.Printf("Site loaded: %s, StatusCode: %d \n", res.Url, res.StatusCode)
	}

	return nil
}

//Make request returning with the custom object with more informations.
func MakeRequestWithDetails(url string) error {
	res, err := helpers.Requester(url)

	if err != nil {
		return err
	}

	if res.StatusCode == 200 {
		//fmt.Printf("Site loaded %s \n your status is %d \n timestamp: %s \n body: %s", res.Url, res.StatusCode, res.Timestamp, string(*res.Body))

		file := fmt.Sprintf("data_%s.html", string(res.Url.Host))

		f, err := os.Create("./data/" + file)
		if err != nil {
			log.Panicf("Error in creating the file of %s\n Error: %s", f.Name(), err)
		}

		defer f.Close()

		f.WriteString(string(*res.Body))

		fmt.Printf("Site loaded %s \n your status is %d \n timestamp: %s \n file: %s \n", res.Url, res.StatusCode, res.Timestamp, f.Name())

	}
	return nil
}
