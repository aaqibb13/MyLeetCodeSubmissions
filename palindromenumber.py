class Solution:
    def isPalindrome(self, x: int) -> bool:
        rev_x = 0
        x_n = x
        pos_limit = 0x7fffffff
        if x < 0:
            return False
        elif x > 0:
            while x != 0:
                rem = (x % 10)
                rev_x = (rev_x * 10) + rem
                x = x//10
            if rev_x & pos_limit == rev_x:
                if x_n == rev_x:
                    return True
                else:
                    return False
        else:
            return True
