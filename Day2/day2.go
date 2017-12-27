package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func checksum(data []string) {
	lines := len(data) - 1
	sum := 0
	for i, _ := range data {
		if i == lines {
			break
		}
		nums := strings.Fields(data[i])
		min, _ := strconv.Atoi(nums[0])
		max := min
		for j, _ := range nums {
			temp, _ := strconv.Atoi(nums[j])
			if temp > max {
				max = temp
			}
			if temp < min {
				min = temp
			}
		}
		sum += (max - min)
	}
	fmt.Printf("sum:%d\n", sum)
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")

	checksum(lines)
}
