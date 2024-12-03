package main

import (
	"bufio"
	"fmt"
	"os"
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

	set := make([][]int, 1000)

	i := 0
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		set[i] = make([]int, len(s))
		for j, v := range s {
			set[i][j], _ = strconv.Atoi(v)
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	reports := 0
	for _, v := range set {
		i1, i2 := v[0], v[1]

		increasing := true

		if i1 > i2 {
			increasing = false
		}

		ok := true
		ignoreable := true

		if increasing {
			for j := 1; j < len(v); j++ {
				fmt.Print(v[j], " ")
				if v[j]-v[j-1] > 3 || v[j]-v[j-1] < 1 {
					if ignoreable {
						fmt.Print("Ignoring ")
						ignoreable = false
						continue
					}

					fmt.Println(v, "Not OK increasing")
					ok = false
					break
				}
			}
			if ok {
				fmt.Println(v, "OK increasing")
				reports++
			}
		} else {
			for j := 1; j < len(v); j++ {
				if ignoreable {
					fmt.Print("Ignoring ")
					ignoreable = false
					continue
				}
				fmt.Print(v[j], " ")
				if v[j-1]-v[j] > 3 || v[j-1]-v[j] < 1 {
					fmt.Println(v, "Not OK decreasing")
					ok = false
					break
				}
			}
			if ok {
				fmt.Println(v, "OK decreasing")
				reports++
			}
		}
	}
	fmt.Println("Reports:", reports)

}
