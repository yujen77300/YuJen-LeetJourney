package besttimetobuyandsellstock

func maxProfit(prices []int) int {

	maxProfit := 0
	buyDay := 0
	sellDay := 1

	for sellDay < len(prices) {
		if prices[buyDay] < prices[sellDay] {
			maxProfit = max(maxProfit, prices[sellDay]-prices[buyDay])
		} else {
			buyDay = sellDay
		}
		sellDay++
	}

	return maxProfit

}
