package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	data, err := os.ReadFile("./input")
	check(err)

	allCalories := strings.Split(string(data), "\n")
	// fmt.Println(allCalories[0])
	current := 0
	highest := 0

	for _, value := range allCalories {
		if value != "" {
			value, _ := strconv.Atoi(value)
			current += value
		} else {
			if current >= highest {
				highest = current
			}
			current = 0
		}
	}

	fmt.Print(highest)
}
