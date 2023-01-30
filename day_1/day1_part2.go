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

	topThreeSum := 0

	elvesCountClone := make([]int, len(elvesCount))
	copy(elvesCountClone, elvesCount)

	for i := 0; i < 3; i++ {
		max := elvesCountClone[0]
		index := 0

		for i, _ := range elvesCountClone {
			if elvesCountClone[i] > max {
				max = elvesCountClone[i]
				index = i
			}
		}

		topThreeSum += max

		if index == 0 {
			elvesCountClone = elvesCountClone[1:]
		} else if index == len(elvesCountClone) {
			elvesCountClone = elvesCountClone[:len(elvesCountClone)-1]
		} else {
			elvesCountClone = append(elvesCountClone[0:index], elvesCountClone[index+1:]...)
		}

	}

	print(topThreeSum)
}
