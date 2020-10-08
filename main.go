package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	var numbers []int
	var sum int = 0

	// Read from console
	fmt.Scan(&input)

	// Scan values from input string
	N, _ := strconv.Atoi(input)

	for i := 0; i < N; i++ {
		fmt.Scan(&input)
		number, _ := strconv.Atoi(input)
		numbers = append(numbers, number)
	}

	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println(sum)
}
