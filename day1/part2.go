package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	ptr := 50
	password := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		splitArray := strings.Split(scanner.Text(), "")

		firstLetter := splitArray[0]

		removeFirstLetter := splitArray[1:]
		combinedString := strings.Join(removeFirstLetter, "")

		moveLockNum, err := strconv.Atoi(combinedString)
		if err != nil {
			fmt.Printf("Error converting string to integer: %v\n", err)
			return
		}

		if firstLetter == "R" {
			for i := 0; i < moveLockNum; i++ {
				ptr = ptr + 1
				ptr = absMod(ptr, 100)
				if ptr == 0 {
					password += 1
				}
			}
		} else if firstLetter == "L" {
			for i := 0; i < moveLockNum; i++ {
				ptr = ptr - 1
				ptr = absMod(ptr, 100)
				if ptr == 0 {
					password += 1
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(password)

}

func absMod(a, b int) int {
	return (a%b + b) % b
}
