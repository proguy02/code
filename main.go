package main

import (
	"code/utils"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}
	input := os.Args[1]
	output := os.Args[2]

	text, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("error reading file", err)
		return
	}
	words := string(text)
	words = utils.String(words)
	words = utils.AnA(words)
	words = utils.Qoute(words)
	words = utils.Base(words)
	words = utils.Puncts(words)

	err = os.WriteFile(output, []byte(words), 0644)
	if err != nil {
		fmt.Println("error writing to file", err)
		return
	}
	fmt.Println("file modified successfully.")
}
