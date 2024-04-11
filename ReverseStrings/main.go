package main

import (
	"fmt"
	"slices"
	"strings"
)

func ReversoStrings(s string) string {
	stringSlice := strings.Split(s, "")
	slices.Reverse(stringSlice)
	return strings.Join(stringSlice, "")
}

func ReverseStringso(s string) string {
	var sample string
	for i := len(s) - 1; i >= 0; i-- {
		sample += string(s[i])
	}
	return sample
}

func ReverseStringBuilder(s string) string {
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}

func main() {
	fmt.Println("Reverse strings implementation")
	fmt.Println(ReversoStrings("Deez"))
	fmt.Println(ReverseStringso("zeed"))
}
