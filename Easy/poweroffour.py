class Solution:
    def isPowerOfFour(self, n: int) -> bool:
        if n==0:
            return False
        if n==1:
            return True
        if (((n%10==4) or (n%10==6)) and ((n&(n-1))==0)):
            return True
        else:
            return False
