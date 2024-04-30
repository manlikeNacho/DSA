package main

import "fmt"

func numberOfSteps(num int) int {
	var numOfSteps int

	for num != 0 {
		if num%2 == 0 {
			num = num / 2
			numOfSteps++
		}

		if num%2 != 0 {
			num--
			numOfSteps++
		}
	}

	return numOfSteps
}

func main() {
	fmt.Println(numberOfSteps(14))
}
