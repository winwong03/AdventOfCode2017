package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func numPass(data []string) {
	count := 0
	numData := len(data) - 1

	for i := 0; i < numData; i++ {
		dupe := false
		wordList := strings.Fields(data[i])
		wordMap := make(map[string]int)
		for _, value := range wordList {
			if wordMap[value] > 0 {
				dupe = true
				break
			}
			wordMap[value] += 1
		}
		wordMap = nil
		if dupe == false {
			count++
		}
	}
	fmt.Printf("Count: %d\n", count)
}

func mergeString(left, right string) string {
	leftLength := len(left)
	rightLength := len(right)

	if leftLength == 0 {
		return right
	}

	if rightLength == 0 {
		return left
	}

	result := make([]string, leftLength+rightLength)

	i, j, k := 0, 0, 0
	for ok := true; ok; ok = (i < leftLength && j < rightLength) {
		if left[i] < right[j] {
			result[k] = string(left[i])
			k++
			i++
		} else {
			result[k] = string(right[j])
			k++
			j++
		}
	}
	res := strings.Join(result, "")
	if k < leftLength+rightLength {
		if i > leftLength-1 {
			res += string([]rune(right[j:rightLength]))

		} else {
			res += string([]rune(left[i:leftLength]))
		}
	}
	return res

}

func sortWord(word string) string {
	length := len(word)
	if length <= 1 {
		return word
	}
	var rightList, leftList string
	if length%2 == 0 {
		rightList = sortWord(string([]rune(word[0 : length/2])))
		leftList = sortWord(string([]rune(word[length/2 : length])))
	} else {
		rightList = sortWord(string([]rune(word[0 : length/2+1])))
		leftList = sortWord(string([]rune(word[length/2+1 : length])))
	}
	return mergeString(rightList, leftList)
}

func noAnagrams(data []string) {
	count := 0
	numData := len(data) - 1

	for i := 0; i < numData; i++ {
		dupe := false
		wordList := strings.Fields(data[i])
		wordMap := make(map[string]int)
		for _, value := range wordList {
			sorted := sortWord(value)
			if wordMap[sorted] > 0 {
				dupe = true
				break
			}
			wordMap[sorted] += 1
		}
		wordMap = nil
		if dupe == false {
			count++
		}
	}
	fmt.Printf("No Anagram Count: %d\n", count)
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("error reading file", err)
	}
	dataList := strings.Split(string(input), "\n")
	numPass(dataList)
	noAnagrams(dataList)

}
