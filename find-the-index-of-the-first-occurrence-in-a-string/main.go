package main

import "fmt"

func strStr(haystack string, needle string) int {

	for i := 0; i < len(haystack); i++ {
		if i+len(needle) > len(haystack) {
			return -1
		} else if haystack[i:i+len(needle)] == needle {

			return i

		}
	}

	return -1
}

func main() {
	//s := "sadbutsad"
	//fmt.Println(s[3:6])
	fmt.Println(strStr("mississippi", "issi"))
}
