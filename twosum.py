# Calculating the sum of two elements in a list such that their sum is equal to the target element

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        h = {}
        for i, num in enumerate(nums):
            n = target-num
            if n not in h:
                h[num] = i
            else:
                return [h[n], i]
