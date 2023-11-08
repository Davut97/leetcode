package main

import "fmt"

func searchInsert(nums []int, target int) int {

	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}

	}

	return left
}
func main() {
	xi := []int{1, 3, 5, 6}
	//fmt.Println(len(xi))
	fmt.Println(searchInsert(xi, 2))

}
