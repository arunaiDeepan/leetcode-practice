class Solution:
    def isGoodString(self, s):
        
        for i in range(len(s) - 1):
            ord_res = abs(ord(s[i+1]) - ord(s[i]))
            
            if ord_res == 1 or ord_res == 25:
                continue
            else:
                return "NO"

        return "YES"