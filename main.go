package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatalf("Error opening file %v", err)
	}
	for line := range getLinesChannel(file) {
		fmt.Printf("read: %s\n", line)
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	linesChan := make(chan string)

	buffer := make([]byte, 8)
	currentLine := ""

	go func() {
		defer f.Close()
		for {
			n, err := f.Read(buffer)
			if err != nil {
				if err == io.EOF {
					close(linesChan)
					break
				}
				log.Fatalf("Error reading file %v", err)
			}
			parts := strings.Split(string(buffer[:n]), "\n")
			currentLine += strings.Join(parts[:1], " ")
			if len(parts) > 1 {
				linesChan <- currentLine
				currentLine = ""
				currentLine += string(parts[1])
			}
		}
	}()
	return linesChan
}
