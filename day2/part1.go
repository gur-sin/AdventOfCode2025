package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkRepeat(start int, end int) []int {
	invalidIDs := []int{}

	for num := start; num <= end; num++ {
		s := strconv.Itoa(num)

		for i := 1; i < len(s); i++ {
			part1 := s[:i]
			part2 := s[i:]

			if part1 == part2 {
				invalidIDs = append(invalidIDs, num)
				break
			}
		}
	}

	return invalidIDs
}

func main() {
	totalInvalidIDs := []int{}
	sum := 0

	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	fileContentInBytes, err := io.ReadAll(reader)
	check(err)
	fileContent := string(fileContentInBytes)
	fileContent = strings.TrimSuffix(fileContent, "\n")

	splitArray := strings.Split(fileContent, ",")
	for _, idRange := range splitArray {
		splitRange := strings.Split(idRange, "-")
		start, err := strconv.Atoi(splitRange[0])
		check(err)
		end, err := strconv.Atoi(splitRange[1])
		check(err)
		invalidIDsInRange := checkRepeat(start, end)

		for _, id := range invalidIDsInRange {
			totalInvalidIDs = append(totalInvalidIDs, id)
		}
	}

	for _, id := range totalInvalidIDs {
		sum += id
	}

	fmt.Println(sum)
}
