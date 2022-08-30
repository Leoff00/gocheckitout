package main

import (
	"fmt"
	"os"

	"github.com/leoff00/gocheckitout/api"
	"github.com/leoff00/gocheckitout/helpers"
)

func firstLog() {
	fmt.Println("Welcome to go Check-it-out, Choose an option to execute the features")
}

func readOptionCode() int {
	var code int
	fmt.Scan(&code)
	fmt.Println("The code was", code)
	return code
}

func main() {
	firstLog()
	urls := helpers.ReadFiles("urls.txt")
	readedCode := readOptionCode()

	switch readedCode {
	case 1:
		for _, url := range urls {
			api.MakeRequest(url)
		}
	case 2:
		for _, url := range urls {
			api.MakeRequestWithDetails(url)
		}
	case 3:
		fmt.Println("Exiting...")
		os.Exit(0)

	default:
		fmt.Println("Invalid operation, exiting...")
		os.Exit(-1)
	}

}
