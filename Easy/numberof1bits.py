# Compute the number of 1's in an unsigned integer
class Solution:
    def hammingWeight(self, n: int) -> int:
        count = 0
        new = str(bin(n))
        for i in new:
            if i == '1':
                count += 1
        return count
