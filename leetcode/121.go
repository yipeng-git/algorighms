package leetcode

func maxProfit(prices []int) int {
	buy := prices[0]
	profit := 0
	for i, _ := range prices {
		newProfit := prices[i] - buy
		if newProfit > profit {
			profit = newProfit
		}
		if prices[i] < buy {
			buy = prices[i]
		}
	}
	return profit
}
