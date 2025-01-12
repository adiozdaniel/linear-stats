package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("😕 Wrong Usage, check documentation")
		return
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Oops, we couldn't locate your file 🥺\nCheck your path and try again\n")
		return
	}
	defer file.Close()

	contents, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("😱 Error reading file...")
		return
	}
	fmt.Printf("File contents: %s\n", contents)
}
