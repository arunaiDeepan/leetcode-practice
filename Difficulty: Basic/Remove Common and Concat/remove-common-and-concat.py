class Solution:
    def concatenatedString(self,s1,s2):
        set1 = set(s1)
        set2 = set(s2)
        
        result = []
        
        for c in s1:
            if c not in set2:
                result.append(c)
                
        for c in s2:
            if c not in set1:
                result.append(c)
                
        return "".join(result) if result else "-1"