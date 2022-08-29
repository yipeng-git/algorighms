package leetcode

func longestConsecutive(nums []int) int {
	cnt := make(map[int]bool, len(nums))
	for i, _ := range nums {
		cnt[nums[i]] = true
	}
	res := 0
	for k, _ := range cnt {
		if _, ok := cnt[k-1]; !ok { //没有当前元素-1
			currLen := 1
			curr := k
			for cnt[curr+1] {
				curr = curr + 1
				currLen = currLen + 1
			}
			if currLen > res {
				res = currLen
			}
		}
	}
	return res
}
