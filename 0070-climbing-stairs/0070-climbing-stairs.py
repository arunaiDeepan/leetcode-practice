class Solution(object):
    def climbStairs(self, n):
        """ :type n: int :rtype: int """
        prev = 1
        prev_ = 0
        current = 0
        for i in range(0,n):
            current = prev + prev_
            prev_ = prev
            prev = current
        return prev 