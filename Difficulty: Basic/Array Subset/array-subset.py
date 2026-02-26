class Solution:
    #Function to check if a is a subset of b.
    def isSubset(self, a, b):
        freq = {}
        
        for i in a:
            freq[i] = freq.get(i,0) + 1
        
        for j in b:
            if j not in freq:
                return False
            
            freq[j] -= 1
            
            if freq[j] < 0:
                return False
                
        return True
