package solution

func characterReplacement(s string, k int) int {
	hitMap := make(map[byte]int)
	left := 0
	count := 0
	result := 0
	for right := 0; right < len(s); right++ {
		hitMap[s[right]]++
		if hitMap[s[right]] > count {
			count = hitMap[s[right]]
		}
		for right-left+1-count > k {
			hitMap[s[left]]--
			left++
		}
		if right-left+1 > result {
			result = right - left + 1
		}
	}
	return result
}
