
from typing import List


class Solution:
    def minTime(self, N : int, jobs : List[int]) -> int:
        
        jobs = sorted(jobs)
        
        def can_be_done(T):
            
            j_id = 0 ## current job pointer
            machines_used = 0
            
            while j_id < N and machines_used < 3:
                
                machines_used += 1 
                
                limit = jobs[j_id] + 2*T
                
                while j_id < N and jobs[j_id] <= limit:
                    j_id += 1
                 
            return j_id >= N
            
        l_t, h_t = 0, max(jobs) - min(jobs)
        
        res = h_t
        
        while l_t <= h_t:
            
            mid = (l_t + h_t) // 2
            
            if can_be_done(mid):
                res = mid 
                h_t = mid - 1
            else:
                l_t = mid + 1
                
        return res