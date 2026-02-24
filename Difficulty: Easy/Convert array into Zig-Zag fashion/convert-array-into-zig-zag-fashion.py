
from typing import List


class Solution:
    def zigZag(self,arr : List[int]) -> None:
        flag = True
        for i in range(len(arr)-1):
            if flag:
               if arr[i+1] < arr[i]:
                   arr[i+1], arr[i] = arr[i], arr[i+1]
            else:
                if arr[i+1] > arr[i]:
                    arr[i], arr[i+1] = arr[i+1], arr[i]
                    
            flag = bool(1-flag)
            
        
