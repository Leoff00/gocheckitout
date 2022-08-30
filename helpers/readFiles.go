package helpers

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func ReadFiles(path string) []string {
	var urls []string
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	buffer := bufio.NewReader(file)

	for {
		line, err := buffer.ReadString('\n')
		line = strings.TrimSpace(line)
		urls = append(urls, line)

		if err == io.EOF {
			break
		}
	}

	return urls

}
