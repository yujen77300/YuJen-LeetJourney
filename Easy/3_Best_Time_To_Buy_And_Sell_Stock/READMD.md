# [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) 


## Solution 1: Two Pointers
- **Time Complexity**: O(n)
- **Space Complexity**: O(1)
- **Approach**: Use two pointers to track the potential buy and sell days. While iterating through the prices, if the current price is lower than buy price, update the buy day. Otherwise, calculate the potential profit and update the maximum profit if needed.

```python
def maxProfit(self, prices: List[int]) -> int:
    max_profit = 0
    buy_day = 0
    sell_day = 1

    while sell_day < len(prices):
        if prices[buy_day] < prices[sell_day]:
            max_profit = max(max_profit, prices[sell_day] - prices[buy_day])
        else:
            buy_day = sell_day
        sell_day += 1
    return max_profit
```

```go
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

```