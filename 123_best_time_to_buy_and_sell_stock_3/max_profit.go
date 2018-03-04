func maxProfit(prices []int) int {
    buy1, buy2, sell1, sell2 := ^int(^uint(0)>>1), ^int(^uint(0)>>1), 0, 0
    for i := 0; i < len(prices); i++ {
        sell2 = int(math.Max(float64(buy2 + prices[i]), float64(sell2)))
        buy2 = int(math.Max(float64(sell1 - prices[i]), float64(buy2)))
        sell1 = int(math.Max(float64(buy1 + prices[i]), float64(sell1)))
        buy1 = int(math.Max(float64(-prices[i]), float64(buy1)))
    }
    return sell2
}