class Solution:
    def buildLowestNumber(self, S,N):
        st = []
        
        for v in S:
            while st and N > 0 and st[-1] > v:
                st.pop()
                N -= 1
            
            st.append(v)
            
        while N > 0:
            st.pop()
            N -= 1
                
        result = "".join(st).lstrip('0')
        
        return result if result else "0"