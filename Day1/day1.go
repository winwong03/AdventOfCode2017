package main

import (
	"io/ioutil"
	"log"
	"strconv"
)

func getSum(input string) {
	sum := 0
	// -1 for the null character at the end
	len := len(input) - 1
	for i := range input {
		input1 := input[i]
		if input1 == input[(i+1)%len] {
			num, err := strconv.Atoi(string(input1))
			if err != nil {
				log.Fatal("Could not turn string into int", err)
			}
			sum += num
		}
	}
	log.Print("Sum:", sum)
}

func main() {
	input, err := ioutil.ReadFile("./day1input.txt")
	if err != nil {
		log.Fatal("error reading file", err)
	}

	getSum(string(input))
}
