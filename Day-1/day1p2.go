package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/set"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var sum int
	var isAnswerFound = false
	s := set.New(set.NonThreadSafe)
	s.Add(sum)

	scanner := bufio.NewScanner(file)
	for !isAnswerFound {
		for scanner.Scan() {
			str := scanner.Text()
			num, _ := strconv.Atoi(str)
			sum += num
			if s.Has(sum) {
				fmt.Println("first twice = ", sum)
				isAnswerFound = true
				break
			} else {
				s.Add(sum)
			}
		}
		file.Seek(0, 0)
		scanner = bufio.NewScanner(file)
	}
}
