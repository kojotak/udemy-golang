package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("filename parameter not provided!")
		os.Exit(1)
	}
	fileName := os.Args[1]
	file, openErr := os.Open(fileName)
	if openErr != nil {
		fmt.Println("can not open", fileName, "because", openErr)
	}
	bufferSize := 1024
	buffer := make([]byte, bufferSize)
	for {
		count, readErr := file.Read(buffer)
		if readErr != nil {
			fmt.Println("can not read", fileName, "because", readErr)
		}
		fmt.Print(string(buffer))
		if count < bufferSize {
			break //nothing left to read
		}
	}
	fmt.Println("")
}
