package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

func divisible_checksum(data []string) {
	length := len(data) - 1
	sum := 0

	for i := 0; i < length; i++ {
		//brute force
		nums := strings.Fields(data[i])
		for j, _ := range nums {
			temp1, _ := strconv.Atoi(nums[j])
			done := false
			for k := j + 1; k < len(nums); k++ {
				temp2, _ := strconv.Atoi(nums[k])
				if math.Mod(float64(temp1), float64(temp2)) == 0 || math.Mod(float64(temp2), float64(temp1)) == 0 {
					if temp1 > temp2 {
						sum += temp1 / temp2
					} else {
						sum += temp2 / temp1
					}
					done = true
					break
				}
			}
			if done == true {
				break
			}
		}
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
	divisible_checksum(lines)
}
