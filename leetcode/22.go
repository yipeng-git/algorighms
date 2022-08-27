package leetcode

//func generateParenthesis(n int) []string {
//	if n == 1 {
//		return []string{"()"}
//	}
//	tmp := generateParenthesis(n - 1)
//	m := map[string]bool{}
//	var res []string
//	for i, _ := range tmp {
//		a := "(" + tmp[i] + ")"
//		if _, ok := m[a]; !ok {
//			m[a] = true
//			res = append(res, a)
//		}
//		b := "()" + tmp[i]
//		if _, ok := m[b]; !ok {
//			m[b] = true
//			res = append(res, b)
//		}
//		c := tmp[i] + "()"
//		if _, ok := m[c]; !ok {
//			m[c] = true
//			res = append(res, c)
//		}
//	}
//	return res
//}

func generateParenthesis(n int) []string {
	var res []string
	generateParenthesisHelper(&res, "", 0, 0, n)

	return res
}

func generateParenthesisHelper(res *[]string, curr string, open, close, max int) {
	if len(curr) == max*2 {
		*res = append(*res, curr)
	}
	if open < max {
		generateParenthesisHelper(res, curr+"(", open+1, close, max)
	}
	if open > close {
		generateParenthesisHelper(res, curr+")", open, close+1, max)

	}
}
