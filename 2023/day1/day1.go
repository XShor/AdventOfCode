package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// readFile, err := os.Open("input-example.txt")
	readFile, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var input []string
	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}

	readFile.Close()

	// fmt.Println(input)
	var sum int
	for _, line := range input {
		reg, err := regexp.Compile("[^0-9]+")
		if err != nil {
			panic(err)
		}
		numericStr := reg.ReplaceAllString(line, "")

		calibration1, _ := strconv.Atoi(numericStr[0:1])
		calibration2, _ := strconv.Atoi(numericStr[len(numericStr)-1:])

		sum += calibration1*10 + calibration2
	}
	fmt.Println(sum)
}
