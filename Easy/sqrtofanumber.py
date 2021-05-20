#Square root of a non-negative integer
class Solution:
    def mySqrt(self, x: int) -> int:
      if x>=0:
        sqr = x**0.5
        return floor(sqr)
