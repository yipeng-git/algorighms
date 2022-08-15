package leetcode

func lengthOfLongestSubstring(s string) int {
	cnt := make(map[byte]int)
	res := 0
	start := 0
	for i, _ := range s {
		if cnt[s[i]] > 0 {
			for true {
				cnt[s[start]] = cnt[s[start]] - 1
				if s[start] == s[i] {
					start += 1
					break
				}
				start += 1
			}
		}
		cnt[s[i]] = cnt[s[i]] + 1
		if i-start+1 > res {
			res = i - start + 1
		}
	}
	return res
}
