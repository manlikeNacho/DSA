package main

import "fmt"

func fizzBuzz(num int) {
	for i := 0; i <= num; i++ {
		if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else if i%15 == 0 {
			fmt.Println("fizz buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fmt.Println("FizzBuzz problem...")
	fizzBuzz(5)
}
