func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }
    min, profit := prices[0], 0
    for i := 1; i < len(prices); i++ {
        if prices[i] > min {
            if profit < prices[i] - min {
                profit = prices[i] - min
            }
        } else {
            min = prices[i]
        }
    }
    return profit
}