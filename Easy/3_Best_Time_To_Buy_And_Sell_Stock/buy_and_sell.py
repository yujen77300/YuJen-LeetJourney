from typing import List
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        max_profit = 0
        buy_day = 0
        sell_day = 1

        while sell_day < len(prices):
            if prices[buy_day] < prices[sell_day] :
                max_profit = max(max_profit, prices[sell_day] - prices[buy_day])
            else:
                buy_day = sell_day
            sell_day += 1
        return max_profit
