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
	defer file.Close()

	buffer := make([]byte, 8)
	currentLine := ""

	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error reading file %v", err)
		}
		parts := strings.Split(string(buffer[:n]), "\n")
		currentLine += strings.Join(parts[:1], " ")
		if len(parts) > 1 {
			fmt.Printf("read: %s\n", currentLine)
			currentLine = ""
			currentLine += string(parts[1])
		}
	}
}
