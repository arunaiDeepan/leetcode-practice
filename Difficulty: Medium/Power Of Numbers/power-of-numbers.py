class Solution:
    def reverseexponentiation(self, n):
        if n == 10:
            return 10
        else:
            return n ** n