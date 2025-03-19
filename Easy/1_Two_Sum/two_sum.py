from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        store_dic = {}
        for index, value in enumerate(nums):
            if (num2 := target - value) in store_dic:
                return [store_dic[num2],index]
            store_dic[value] = index
        return []
