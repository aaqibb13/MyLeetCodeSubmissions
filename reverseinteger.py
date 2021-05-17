class Solution:
    def reverse(self, x: int) -> int:
        rev_x = 0
        if x < 0:
            x = abs(x)
            while x != 0:
                rem = (x % 10)
                rev_x = (rev_x * 10) + rem
                x = x//10
            return -rev_x
        elif x > 0:
            while x != 0:
                rem = (x % 10)
                rev_x = (rev_x * 10) + rem
                x = x//10
            return rev_x
        else:
            return 0
