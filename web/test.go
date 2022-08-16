package main

import "fmt"

func main() {
	nums := []int{1, 123, 35642, 5, 7, 87, 213, 3456, 123, 1}
	max := Max(nums, 0, len(nums)-1)
	fmt.Println(max)
}

func Max(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	}
	mid := left + ((right - left) >> 1)
	leftMax := Max(nums, left, mid)
	rightMax := Max(nums, mid+1, right)
	if leftMax > rightMax {
		return leftMax
	} else {
		return rightMax
	}
}
