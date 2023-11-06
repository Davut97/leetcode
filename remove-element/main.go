package main

func removeElement(nums []int, val int) int {

	counter := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[counter] = nums[i]
			counter++
		}

	}

	return counter
}
func main() {
	xi := []int{3, 2, 2, 3}
	removeElement(xi, 3)

}
