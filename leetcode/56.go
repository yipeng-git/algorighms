package leetcode

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return false
	})
	res := [][]int{}
	start := intervals[0][0]
	end := intervals[0][1]
	for i := 1; i < len(intervals); i += 1 {
		if intervals[i][0] <= end {
			if intervals[i][1] > end {
				end = intervals[i][1]
			}
		} else {
			res = append(res, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		}
	}
	res = append(res, []int{start, end})

	return res
}
