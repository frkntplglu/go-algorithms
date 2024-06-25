package main

func scoreOfString(s string) int {
	total := 0

	for i := 0; i < len(s)-1; i++ {
		total += abs(int(s[i]) - int(s[i+1]))
	}

	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
