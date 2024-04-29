package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}

	return fib(n-2) + fib(n-1)
}

func main() {
	fmt.Println(fib(5))
}
