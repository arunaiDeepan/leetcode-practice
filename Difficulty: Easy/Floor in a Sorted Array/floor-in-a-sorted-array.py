class Solution:
    def findFloor(self, arr, x):
        
        floor_i = -1
        
        for i in range(len(arr)):
            if arr[i] <= x:
                floor_i = i

        return floor_i