package points_to_offer

/*
在一个长度为 n 的数组里的所有数字都在 0 到 n-1 的范围内。数组中某些
数字是重复的，但不知道有几个数字是重复的，也不知道每个数字重复几次。请
找出数组中任意一个重复的数字。
*/

func Solution3(slice []int) int {
	pos := 0
	for {
		if pos >= len(slice) {
			return -1
		}
		cur := slice[pos]
		if cur == pos {
			pos += 1
			continue
		}
		next := slice[cur]
		if next == cur {
			return next
		}
		slice[pos] = next
		slice[cur] = cur
		if slice[pos] == pos {
			pos += 1
			continue
		}
	}
}
