package lengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	maxLength := 0
	subMap := make(map[byte]int)
	var i, j int
	for {
		if j == len(s) {
			if length := j - i; maxLength < length {
				maxLength = length
			}
			break
		}
		if i == j {
			subMap[s[i]] = i
			j++
			continue
		}
		if _, ok := subMap[s[j]]; !ok {
			subMap[s[j]] = j
			j++
		} else {
			if length := j - i; maxLength < length {
				maxLength = length
			}
			i = subMap[s[j]] + 1
			j = i
			subMap = make(map[byte]int)
		}
	}
	return maxLength
}
