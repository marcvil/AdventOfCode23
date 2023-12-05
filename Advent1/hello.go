package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var slice []string
	slice = make([]string, 0, 2000)
	scanner := bufio.NewScanner(file)
	// optionall	y, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}

	regexPattern := regexp.MustCompile(`(?:(one|two|three|four|five|six|seven|eight|nine))`)

	values := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var sum int

	for i := 0; i < len(slice); i++ {
		matches := regexPattern.FindAllString(slice[i], -1)

		var first string
		var last string
		var numberString string

		first = getValue(matches[0], values)
		last = getValue(matches[len(matches)-1], values)
		numberString = first + last

		fmt.Println(numberString)

		int1, err := strconv.Atoi(numberString)
		if err != nil {
			log.Fatal(err)
		}
		sum = sum + int1
	}

	log.Printf("Sum is: %d", sum)

}

func getValue(str string, dict map[string]string) string {
	if len(str) > 1 {
		return dict[str]
	}
	return str
}
