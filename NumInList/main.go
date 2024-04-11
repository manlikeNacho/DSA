package main

import "fmt"

func NumInList(list []int, num int) bool {
	for _, val := range list {
		if num == val {
			return true
		}
	}

	return false
}

func SumAllNum(list []int) int {
	sum := 0

	for _, val := range list {
		sum += val
	}

	return sum
}

func SumAllNumRecursive(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	return numbers[0] + SumAllNumRecursive(numbers[1:])
}

func main() {
	fmt.Println("Number in List")
	var numbers = []int{10, 20, 30, 40, -50}
	fmt.Println(NumInList(numbers, 20))
	fmt.Println(SumAllNum(numbers))
	fmt.Println(SumAllNumRecursive(numbers))
}
