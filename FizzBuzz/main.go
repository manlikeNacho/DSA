package main

import "fmt"

func fizzBuzz(num int) int {
	if num == 0 {
		return 0
	}
	// for i := 0; i <= num; i++ {
	// 	if i%3 == 0 {
	// 		fmt.Println("fizz")
	// 	} else if i%5 == 0 {
	// 		fmt.Println("buzz")
	// 	} else if i%15 == 0 {
	// 		fmt.Println("fizz buzz")
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }

	return num + fizzBuzz(num-1)
}

func main() {
	fmt.Println("FizzBuzz problem...")
	fmt.Println(fizzBuzz(10))
}
