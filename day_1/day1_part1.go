package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	readFile, err := os.Open("day1_elfCalories.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	elvesCount := []int{}
	count := 0

	for fileScanner.Scan() {

		if fileScanner.Text() == "" {
			elvesCount = append(elvesCount, count)
			count = 0
			continue
		}
		value, _ := strconv.Atoi(fileScanner.Text())
		count += value
	}

	max := elvesCount[0]

	for i, _ := range elvesCount {
		if elvesCount[i] > max {
			max = elvesCount[i]
		}
	}

	print(max)
}
