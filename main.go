package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR")
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	pieces, err := parseTetrominoes(string(data))
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	for _, row := range solve(pieces) {
		fmt.Println(string(row))
	}
}
