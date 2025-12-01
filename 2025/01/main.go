package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func readFile() string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)

	if err != nil {
		log.Fatalf("failed reading all: %s", err)
	}

	body := string(data)

	return body
}

func main() {

	//dial := 50

	fmt.Println(readFile())

	//for i := 1; i <= 5; i++ {
	//
	//	fmt.Println("i =", 100/i)
	//}
}
