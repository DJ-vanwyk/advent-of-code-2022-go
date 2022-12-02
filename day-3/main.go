package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := input()

	fmt.Println(input)

}

func input() []string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	return lines
}
