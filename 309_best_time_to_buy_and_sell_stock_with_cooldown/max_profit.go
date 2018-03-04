func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }
    preBuy, preSell, buy, sell := 0, 0, ^int(^uint(0) >> 1), 0
    for _, price := range prices {
        preBuy = buy
        if buy < preSell - price {
            buy = preSell - price
        }
        preSell = sell
        if sell < preBuy + price {
            sell = preBuy + price
        }
    }
    return sell
}