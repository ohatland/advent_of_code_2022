package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	start := time.Now()

	data, err := os.ReadFile("./input.txt")
	check(err)

	allCalories := strings.Split(string(data), "\n")
	// fmt.Println(allCalories[0])

	current := 0
	sumCaloriesPerElves := []int{}

	for _, value := range allCalories {
		if value != "" {

			value, _ := strconv.Atoi(value)
			current += value
		} else {

			sumCaloriesPerElves = append(sumCaloriesPerElves, current)
			current = 0
		}
	}
	// fmt.Println(sumCaloriesPerElves)
	sort.Ints(sumCaloriesPerElves[:])

	topTree := 0
	for _, v := range sumCaloriesPerElves[len(sumCaloriesPerElves)-3:] {
		topTree += v
	}
	fmt.Println(topTree)

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
