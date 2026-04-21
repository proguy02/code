// Author Raymond Nicholas. @raymondproguy

package main

import (
	"bufio"
	"fmt"
	"goreloaded/processors"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . sample.txt result.txt")
		return
	}
	input := os.Args[1]
	output := os.Args[2]

	data, err := os.Open(input)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	content := bufio.NewScanner(data)
	var write strings.Builder

	for content.Scan() {
		text := content.Text()

		modified := processors.AnA(text)
		modified = processors.Hex(modified)
		modified = processors.Bin(modified)
		modified = processors.Up(modified)
		modified = processors.Low(modified)
		modified = processors.Cap(modified)
		modified = processors.UpN(modified)
		modified = processors.LowN(modified)
		modified = processors.CapN(modified)
		modified = processors.Puncts(modified)
		modified = processors.Qoute(modified)
		modified = processors.Qoute2(modified)

		write.WriteString(modified)
		write.WriteString("\n")
	}

	res := write.String()
	err = os.WriteFile(output, []byte(res), 0644)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("File modified successfully, check result.txt")
}
