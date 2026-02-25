class Solution:
    def computeParity(self, N):
        
        num_1 = 0

        while N:
            num_1 += N & 1
            N >>= 1
        
        if num_1 % 2 == 0:
            return "even"
        else:
            return "odd"
            