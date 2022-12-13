package helpers

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func ReadFiles(path string, rChan chan []string) chan []string {
	var urls []string
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	buffer := bufio.NewReader(file)

	for {
		line, err := buffer.ReadString('\n')
		line = strings.TrimSpace(line)
		urls = append(urls, line)

		if err == io.EOF {
			break
		}
	}

	rChan <- urls
	return rChan
}
