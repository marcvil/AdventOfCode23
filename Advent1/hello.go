package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var slice []string
	slice = make([]string, 0, 200)
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}

	var sum int

	for i := 0; i < len(slice); i++ {
		var numberString string
		word := []rune(slice[i])
		for j := 0; j < len(word); j++ {
			if unicode.IsDigit(word[j]) {
				log.Printf(string(word[j]))
				numberString += string(word[j])
				break
			}
		}

		for j := len(word) - 1; j >= 0; j-- {
			if unicode.IsDigit(word[j]) {
				log.Printf(string(word[j]))
				numberString += string(word[j])
				break
			}
		}

		int1, err := strconv.Atoi(numberString)
		if err != nil {
			log.Fatal(err)
		}
		sum = sum + int1
	}

	log.Printf("Sum is: %d", sum)
}
