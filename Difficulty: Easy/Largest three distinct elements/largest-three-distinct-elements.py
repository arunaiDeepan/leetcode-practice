class Solution:
    def getThreeLargest(self, arr):
        one = two = th = float('-inf')
        
        for i in arr:
            if i > one:
                th = two
                two = one 
                one = i
            elif i > two and i != one:
                th = two                
                two = i
            elif i > th and i != two and i != one:
                th = i
        
        res = []
        
        if one == float('-inf'):
            return res
        res.append(one)
        
        if two == float('-inf'):
            return res
        res.append(two)
        
        if th == float('-inf'):
            return res
        res.append(th)
        
        return res