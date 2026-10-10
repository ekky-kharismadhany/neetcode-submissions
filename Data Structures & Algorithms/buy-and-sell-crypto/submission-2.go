func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var minPrice int = prices[0]
	var maxProfit = 0

	for _, price := range prices {
		if minPrice > price {
			minPrice = price
		} else if price - minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}
