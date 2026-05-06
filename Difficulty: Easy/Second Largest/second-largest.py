class Solution:
    def getSecondLargest(self, arr):
        n = len(arr) - 1
        
        arr.sort()
        
        for i in range(n-1, -1, -1):
            if arr[i] != arr[n]:
                return arr[i]
                
        return -1