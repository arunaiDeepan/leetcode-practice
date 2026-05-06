class Solution:
    def getAlternates(self, arr):
        res = []
        for i in range(len(arr)):
            if i % 2 == 0:
                res.append(arr[i])
              
        return res