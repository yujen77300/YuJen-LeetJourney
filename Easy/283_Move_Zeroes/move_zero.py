class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        cur = 0
        for i in range(len(nums)):
            if nums[i] != 0 :
                nums[cur] = nums[i]
                cur += 1

        while cur < len(nums):
            nums[cur] = 0
            cur += 1