package main

func removeDuplicates(nums []int) int {
	lenNum := len(nums)
	counter := 0

	for i := 1; i < lenNum; i++ {
		if nums[counter] != nums[i] {

			counter++
			nums[counter] = nums[i]
		}
	}
	return counter + 1
}
func main() {
	x := []int{1, 1, 2, 3, 4}
	removeDuplicates(x)

}
