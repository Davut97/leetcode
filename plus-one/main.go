package main

import "fmt"

func plusOne(digits []int) []int {

	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	xi := make([]int, n+1)
	xi[0] = 1
	return xi
}
func main() {

	fmt.Println(plusOne([]int{9, 9, 9, 9}))
}
