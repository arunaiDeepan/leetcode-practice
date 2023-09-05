class Solution(object):
    def reverse(self, x):
        """ :type x: int :rtype: int """
        INT_MAX = 2**31 - 1
        INT_MIN = -2**31
        rev = 0
        is_negative = False
        
        if x < 0:
            is_negative = True
            x = -x
            
        while x != 0:
        # Get the last digit of x
            pop = x % 10
            x //= 10
            if rev > (INT_MAX - pop) // 10 or rev < (INT_MIN - pop) // 10:
                return 0
            rev = rev * 10 + pop
        if is_negative:
            rev = -rev
        return rev




