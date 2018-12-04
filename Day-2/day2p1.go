package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var twoNum, threeNum = 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		var list [26]int
		var isTwo, isThree = false, false
		for i := 0; i < len(str); i++ {
			char := str[i]
			list[char-97]++
		}
		for i := 0; i < len(list); i++ {
			if list[i] == 2 {
				isTwo = true
			} else if list[i] == 3 {
				isThree = true
			}
		}
		if isTwo {
			twoNum++
		}
		if isThree {
			threeNum++
		}
	}
	fmt.Println("Answer = ", twoNum*threeNum)
}
