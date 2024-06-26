# Reversing an Integer with memory constraint as 32-bit max_int = (2^32 - 1) and min_int = -(2^31)

class Solution:
    def reverse(self, x: int) -> int:
        rev_x = 0
#         neg_limit= -0x80000000
        pos_limit= 0x7fffffff
        if x < 0:
            x = abs(x)
            while x != 0:
                rem = (x % 10)
                rev_x = (rev_x * 10) + rem
                x = x//10
            if rev_x & pos_limit == rev_x:
                return -rev_x
            else:
                return 0
        elif x > 0:
            while x != 0:
                rem = (x % 10)
                rev_x = (rev_x * 10) + rem
                x = x//10
            if rev_x & pos_limit == rev_x:
                return rev_x
            else:
                return 0
        else:
            return 0
