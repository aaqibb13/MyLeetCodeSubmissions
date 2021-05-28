class Solution:
    def isPowerOfTwo(self, n: int):
        global count
        if n < 0:
            return False
        else:
            num = bin(n)
            count = 0
            for i in num:
                if i == '1':
                    count += 1

            if count == 1:
                return True
            else:
                return False
