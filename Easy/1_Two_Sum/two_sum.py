class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        store_dic = {}
        for index, value in enumerate(nums):
            num1 = value
            num2 = target - num1
            if num2 in store_dic:
                index2 = store_dic[num2]
                return [index, index2]
            else:
                store_dic[num1] = index
        return []
