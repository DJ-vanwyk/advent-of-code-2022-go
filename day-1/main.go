package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// content, err := os.ReadFile("input.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(content))

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	elfGroups := groupElfs(f)

	total := 0

	for i := 0; i < 3; i++ {
		val, index := maxInt(elfGroups)
		total += val
		elfGroups = remove(elfGroups, index)
	}

	fmt.Println(total)

}

func groupElfs(f *os.File) []int {

	scanner := bufio.NewScanner(f)

	var elves []int

	totalElfCalories := 0

	for scanner.Scan() {

		line := scanner.Text()

		if line == "" {
			elves = append(elves, totalElfCalories)
			totalElfCalories = 0
		} else {
			lineInt, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}

			totalElfCalories += lineInt
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return elves
}

func maxInt(slice []int) (int, int) {

	largestNum := 0
	largestNumIndex := -1

	for i, iVal := range slice {
		largestNum = iVal
		largestNumIndex = i
		for _, jVal := range slice {
			if iVal < jVal {
				largestNum = 0
				largestNumIndex = -1
			}
		}

		if largestNum == iVal {
			break
		}

	}

	return largestNum, largestNumIndex
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
