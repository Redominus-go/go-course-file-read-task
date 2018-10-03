package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1]

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	text := make([]byte, 9999)
	_, err = f.Read(text)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(text))
}
