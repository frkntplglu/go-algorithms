package main

func dominantIndex(nums []int) int {
	var maxIndex = 0

	for i, val := range nums {
		if val > nums[maxIndex] {
			maxIndex = i
		}
	}

	for i := range nums {
		if maxIndex != i && nums[maxIndex] < 2*nums[i] {
			return -1
		}
	}

	return maxIndex
}
