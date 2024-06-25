package main

func strStr(haystack string, needle string) int {
	m := len(haystack)
	n := len(needle)

	for i := 0; i <= m-n; i++ {
		var j int
		for j = 0; j < n; j++ {
			if string(haystack[i+j]) != string(needle[j]) {
				break
			}
		}

		if j == n {
			return i
		}
	}
	return -1
}
