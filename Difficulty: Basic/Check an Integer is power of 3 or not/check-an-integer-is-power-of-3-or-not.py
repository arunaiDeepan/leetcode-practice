class Solution:
    def isPowerof3 (ob,N):
        
        while N % 3 == 0:
            N = N // 3
            
        if N == 1:
            return "Yes"
        else:
            return "No"
            