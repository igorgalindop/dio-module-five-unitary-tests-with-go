package main

import "fmt"

func main() {
	resultSum := sum(1, 2, 3)
	resultMulti := multiply(10, 10)
	resultSub := subtract(5, 10)
	resultDiv := divide(20)

	fmt.Println(resultSum, resultMulti, resultSub, resultDiv)
}

func sum(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func multiply(numbers ...int) int {
	total := 1
	for _, value := range numbers {
		total = total * value
	}
	return total
}

func subtract(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		if total == 0 {
			total = value
		} else {
			total = total - value
		}
	}
	return total
}

func divide(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		if total == 0 {
			total = value
		} else {
			total = total / value
		}
	}
	return total
}
