package main

import "fmt"

// s = "MCMXCIV"
// Output: 1994
// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
func romanToInt(s string) int {
	h := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	prevValue := 0
	for i := len(s) - 1; i >= 0; i-- {
		currentValue := h[s[i]]
		if currentValue < prevValue {
			result -= currentValue
		} else {
			result += currentValue
		}
		prevValue = currentValue
	}

	return result
}
