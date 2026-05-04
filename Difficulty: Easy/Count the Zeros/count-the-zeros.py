class Solution:
    def countZeroes(self, arr):
        
        count_z = 0
        
        for i in arr:
            
            if i == 0:
                count_z += 1
                
        return count_z