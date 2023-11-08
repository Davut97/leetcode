package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	xi := strings.Split(s, " ")
	return len(xi[len(xi)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon    "))
}
