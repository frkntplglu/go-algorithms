package main

// First solution
func pivotIndex(nums []int) int {
	var sumLeft, sumRight []int

	for idx, _ := range nums {
		sumLeftTotal := 0
		sumRightTotal := 0
		for i := idx + 1; i < len(nums); i++ {
			sumLeftTotal += nums[i]
		}
		sumLeft = append(sumLeft, sumLeftTotal)
		for j := 0; j < idx; j++ {
			sumRightTotal += nums[j]
		}
		sumRight = append(sumRight, sumRightTotal)

		if sumLeftTotal == sumRightTotal {
			return idx
		}
	}

	return -1
}

// More optimized solution
func pivotIndex(nums []int) int {
	var total = 0
	for i := range nums {
		total += nums[i]
	}

	var cur = 0

	for i := range nums {
		if cur == total-cur-nums[i] {
			return i
		}

		cur += nums[i]
	}

	return -1
}
