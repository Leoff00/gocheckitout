package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/leoff00/gocheckitout/api"
	"github.com/leoff00/gocheckitout/helpers"
	"github.com/leoff00/gocheckitout/logs"
)

func readOptionCode() int16 {
	var code int16
	fmt.Scan(&code)
	return code
}

var wg sync.WaitGroup

func main() {
	logs.WelcomeLog()

	res := make(chan []string)
	go helpers.ReadFiles("dummy.txt", res)
	urls := <-res

	readedCode := readOptionCode()

	switch readedCode {
	case 1:
		for _, url := range urls {
			wg.Add(1)
			go func(nUrl string) {
				err := api.MakeRequest(nUrl)
				if err != nil {
					log.Default().Fatal("Cannot make request correctly.")
				}
				defer wg.Done()
			}(url)
		}
	case 2:
		for _, url := range urls {
			wg.Add(1)
			go func(nUrl string) {
				err := api.MakeRequestWithDetails(nUrl)
				if err != nil {
					log.Default().Fatal("Cannot make request correctly.", err)
				}
				defer wg.Done()
			}(url)
		}
	case 3:
		logs.Exiting()
		os.Exit(0)

	default:
		logs.InvalidOperation()
		os.Exit(-1)
	}
	defer wg.Wait()

}
