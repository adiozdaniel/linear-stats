package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

	data := strings.Split(string(contents), "\n")
	var actualDataY []float64
	var lineX []float64

	for i, line := range data {
		num, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
		if err != nil {
			fmt.Printf("🧐 Invalid data at line: %d:\n`%v`\n", i+1, data[i])
			return
		}

		actualDataY = append(actualDataY, num)
		lineX = append(lineX, float64(i))
	}
	fmt.Printf("actual data: %v\nlineX: %v\n", actualDataY, lineX)
}
