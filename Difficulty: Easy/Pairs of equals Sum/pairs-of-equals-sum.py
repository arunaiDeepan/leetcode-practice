class Solution:
    def findPairs(self, arr): 
        
        n = len(arr)
        
        sum_d = {}
        
        for i in range(n):
            for j in range(i+1, n):
                
                sum_v = arr[i] + arr[j]
                
                if sum_v in sum_d:
                    x = sum_d[sum_v]
                    
                    if x[0] != i and x[0] != j and \
                       x[1] != j and x[1] != j:
                        
                        return True
                
                sum_d[sum_v] = (i,j)
                
        return False