package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := input()

	a(input)
	b(input)

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

func a(input []string) {
	score := 0
	for _, v := range input {
		char := strings.Split(v, " ")
		switch char[1] {
		case "X": //Rock
			switch char[0] {
			case "A": //Rock
				score += 1 + 3
			case "B": //Paper
				score += 1 + 0
			case "C": //Scissors
				score += 1 + 6
			}
		case "Y": //Paper
			switch char[0] {
			case "A": //Rock
				score += 2 + 6
			case "B": //Paper
				score += 2 + 3
			case "C": //Scissors
				score += 2 + 0
			}
		case "Z": //Scissors
			switch char[0] {
			case "A": //Rock
				score += 3 + 0
			case "B": //Paper
				score += 3 + 6
			case "C": //Scissors
				score += 3 + 3
			}

		}

	}

	fmt.Println(score)
}

func b(input []string) {
	score := 0
	for _, v := range input {
		char := strings.Split(v, " ")

		switch char[1] {
		case "X": //Lose
			switch char[0] {
			case "A": //Rock
				score += 0 + 3
			case "B": //Paper
				score += 0 + 1
			case "C": //Scissors
				score += 0 + 2
			}
		case "Y": //Draw
			switch char[0] {
			case "A": //Rock
				score += 3 + 1
			case "B": //Paper
				score += 3 + 2
			case "C": //Scissors
				score += 3 + 3
			}
		case "Z": //Win
			switch char[0] {
			case "A": //Rock
				score += 6 + 2
			case "B": //Paper
				score += 6 + 3
			case "C": //Scissors
				score += 6 + 1
			}

		}

	}

	fmt.Println(score)
}
