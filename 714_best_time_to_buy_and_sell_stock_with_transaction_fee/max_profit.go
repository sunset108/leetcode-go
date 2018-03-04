func maxProfit(prices []int, fee int) int {
	cash, hold := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		cash = int(math.Max(float64(cash), float64(hold + prices[i] - fee)))
		hold = int(math.Max(float64(hold), float64(cash - prices[i])))
}
	return cash
}
