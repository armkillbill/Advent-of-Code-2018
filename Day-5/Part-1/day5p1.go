package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = scanner.Text()
	}

	hasDelete := true
	for hasDelete {
		hasDelete = false
		for i := 0; i < len(input)-1; {
			if input[i]-input[i+1] == 32 || input[i+1]-input[i] == 32 {
				input = input[:i] + input[i+2:]
				hasDelete = true
			} else {
				i++
			}
		}
	}

	fmt.Println("Answer = ", len(input))

}
