# calculate a ** b mod 1337
class Solution:
    def superPow(self, a: int, b: List[int]) -> int:
        s = [str(i) for i in b]
        b_f = int("".join(s))
        res = pow(a,b_f,1337)
        return res
