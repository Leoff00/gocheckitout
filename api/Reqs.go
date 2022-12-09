package api

import (
	"fmt"
	"log"

	"github.com/leoff00/gocheckitout/helpers"
)

func MakeRequest(url string) error {
	fmt.Println("Loading...")
	res, err := helpers.Requester(url)

	if err != nil {
		return err
	}

	if res.StatusCode == 200 {
		fmt.Printf("Site loaded: %s, StatusCode: %d \n", res.Url, res.StatusCode)
	}
	return nil
}

func MakeRequestWithDetails(url string) error {
	fmt.Println("Loading...")
	res, err := helpers.Requester(url)

	if err != nil {
		return err
	}

	if res.StatusCode == 200 {
		fmt.Printf("Site loaded %s \n your status is %d \n timestamp: %s \n", res.Url, res.StatusCode, res.Timestamp.String())
	} else {
		log.Fatal(res.StatusCode, "Error, cannot make a request...")
	}

	return nil
}
