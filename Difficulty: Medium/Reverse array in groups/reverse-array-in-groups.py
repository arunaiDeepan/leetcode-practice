class Solution:
    
    def reverse(self, arr, i, n, k):
        l = i
        r = min(n - 1, i + k-1)
        
        while l < r:
            arr[l], arr[r] = arr[r], arr[l]
            l += 1
            r -= 1

    def reverseingroups(self, arr, k):
        n = len(arr)
        
        for i in range(0, n, k):
            self.reverse(arr, i, n, k)
        
        return arr