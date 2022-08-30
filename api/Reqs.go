package api

import (
	"fmt"
	"log"
	"net/http"
)

func MakeRequest(url string) error {
	fmt.Println("Loading...")
	res, err := http.Get(url)

	if err != nil {
		return err
	}

	if res.StatusCode == 200 {
		fmt.Println("Site loaded:", res.Request.URL, true)
	}
	return nil
}

func MakeRequestWithDetails(url string) error {
	fmt.Println("Loading...")
	res, err := http.Get(url)

	if err != nil {
		return err
	}

	if res.StatusCode == 200 {
		fmt.Println("Site loaded:",
			res.Request.URL, "your status is:",
			res.StatusCode)
	} else {
		log.Fatal(res.StatusCode, "Error, cannot make a request...")
	}

	return nil
}
