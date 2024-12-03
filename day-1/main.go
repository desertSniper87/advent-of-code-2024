package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	set1 := make([]int, 1000)
	set2 := make([]int, 1000)

	i := 0
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		set1[i], err = strconv.Atoi(s[0])
		if err != nil {
			fmt.Printf("Error converting string to int: %v\n", err)
			return
		}
		set2[i], err = strconv.Atoi(s[1])
		if err != nil {
			fmt.Printf("Error converting string to int: %v\n", err)
			return
		}

		i++

	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	sort.Ints(set1)
	sort.Ints(set2)

	sum := 0
	for i, v1 := range set1 {
		v2 := set2[i]

		if v1 > v2 {
			sum += v1 - v2
		} else {
			sum += v2 - v1
		}
	}

	freq := getFrequencyMap(set2)
	sum = 0

	for _, i := range set1 {
		fmt.Println("freq[i]: ", freq[i])
		sum += freq[i] * i
	}

	fmt.Printf("Sum: %d\n", sum)
}

// Using map for multiple elements
func getFrequencyMap(list []int) map[int]int {
	freq := make(map[int]int)
	for _, num := range list {
		freq[num]++
	}
	return freq
}
