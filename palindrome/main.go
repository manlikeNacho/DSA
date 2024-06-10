package main

import "fmt"

func isPalindrome(x int) bool {
	var reverse int = 0

	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	if x == 0 || (x > 0 && x < 10) {
		return true
	}

	for x > reverse {
		//Breaks down the number into sub-divisible parts of 10s
		reverse = x%10 + reverse*10
		fmt.Println(x, reverse)
		//divides the main numbr by 10
		x = x / 10
		fmt.Println(x, reverse)
	}
	return x == reverse || x == reverse/10
}

func main() {

	fmt.Println(12 % 10)
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(11))
}
