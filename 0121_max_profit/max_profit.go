package maxprofit

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	maxProfit, minPrice := 0, prices[0]
	for i := range prices {
		if prices[i] < minPrice {
			minPrice = prices[i]
			continue
		}

		if maxProfit < prices[i]-minPrice {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}
