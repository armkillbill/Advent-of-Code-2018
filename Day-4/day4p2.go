package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var log []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		log = append(log, str)
	}

	sort.Strings(log)

	// for _, v := range log {
	// 	fmt.Println(v)
	// }

	var table [][60]int
	var guardNum, sleep, wake int
	var maxGuardNum = 0
	for _, v := range log {
		var hour, minute int
		var text, date string
		fmt.Sscanf(v, "[%s %d:%d]", &date, &hour, &minute)
		text = v[19:]

		if text == "falls asleep" {
			sleep = minute
		} else if text == "wakes up" {
			wake = minute
			//extend table
			if guardNum > maxGuardNum {
				t := make([][60]int, guardNum+1, guardNum+1)
				copy(t, table)
				table = t
				maxGuardNum = guardNum
			}
			for i := sleep; i < wake; i++ {
				table[guardNum][i]++
			}
		} else {
			fmt.Sscanf(text, "Guard #%d begins shift", &guardNum)
		}
	}

	mostGuardMaxSleep, mostSleepGuardID, minuteOfMostSleep := 0, 0, 0
	for guardID, data := range table {
		minuteOfMaxSleep, guardMaxSleep := 0, 0
		for minute, freq := range data {
			if freq > guardMaxSleep {
				minuteOfMaxSleep = minute
				guardMaxSleep = freq
			}
		}
		if guardMaxSleep > mostGuardMaxSleep {
			mostGuardMaxSleep = guardMaxSleep
			mostSleepGuardID = guardID
			minuteOfMostSleep = minuteOfMaxSleep
		}
	}

	fmt.Println("Answer = ", mostSleepGuardID*minuteOfMostSleep)
}
