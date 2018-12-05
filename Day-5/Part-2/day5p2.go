package main

import (
	"bufio"
	"fmt"
	"math"
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

	minSeq := math.MaxInt32
	for j := 0; j < 26; j++ {
		str := input
		for i := 0; i < len(str); {
			if str[i] == byte(65+j) || str[i] == byte(97+j) {
				str = str[:i] + str[i+1:]
			} else {
				i++
			}
		}
		hasDelete := true
		for hasDelete {
			hasDelete = false
			for i := 0; i < len(str)-1; {
				if str[i]-str[i+1] == 32 || str[i+1]-str[i] == 32 {
					str = str[:i] + str[i+2:]
					hasDelete = true
				} else {
					i++
				}
			}
		}
		if len(str) < minSeq {
			minSeq = len(str)
		}
	}

	fmt.Println("Answer = ", minSeq)
}
