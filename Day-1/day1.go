package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var sum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		num, _ := strconv.Atoi(str)
		sum += num
	}

	fmt.Println("sum = ", sum)
}
