func maxProfit(prices []int) int {
	var minPrice int = math.MaxInt32
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
