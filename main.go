package main

import (
	"fmt"
	"os"

	"github.com/leoff00/gocheckitout/api"
	"github.com/leoff00/gocheckitout/helpers"
	"github.com/leoff00/gocheckitout/logs"
)

func readOptionCode() int {
	var code int
	fmt.Scan(&code)
	return code
}

func main() {

	logs.WelcomeLog()

	res := make(chan []string)
	go helpers.ReadFiles("urls.txt", res)
	urls := <-res
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
		logs.Exiting()
		os.Exit(0)

	default:
		logs.InvalidOperation()
		os.Exit(-1)
	}

}
