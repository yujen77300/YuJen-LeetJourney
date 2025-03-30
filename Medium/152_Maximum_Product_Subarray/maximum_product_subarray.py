from typing import List
class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        min_ending = nums[0]
        max_ending = nums[0]
        maximum_product =  max_ending

        for num in nums[1:]:
            temp = min_ending
            min_ending = min(num, num*min_ending, num*max_ending)
            max_ending = max(num, num*temp, num*max_ending)
            maximum_product = max(max_ending, maximum_product)
        return maximum_product