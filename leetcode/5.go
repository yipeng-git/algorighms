package leetcode

//func longestPalindrome(s string) string {
//	var start int
//	var end int
//	for i := 0; i < len(s); i += 1 {
//		a := longestPalindromeHelper(s, i, i)
//
//		b := longestPalindromeHelper(s, i, i+1)
//		var c int
//		if a > b {
//			c = a
//		} else {
//			c = b
//		}
//		if c > (start - end + 1) {
//
//		}
//	}
//	return res
//}
//
//func longestPalindromeHelper(s string, i, j int) int {
//	for i >= 0 && j < len(s) && s[i] == s[j] {
//		i += 1
//		j += 1
//	}
//	return s[i:j]
//}
