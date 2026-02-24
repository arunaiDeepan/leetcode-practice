class Solution:
    def missingNum(self, arr):
        
        n = len(arr) + 1
        
        t = n * (n+1) // 2
        
        arr_sum = sum(arr)
        
        return t - arr_sum