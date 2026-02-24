def Sandwiched_Vowel(s):
    vowels = {'a','e','i','o','u'}
    
    n = len(s)
    
    result = [s[0]]
    
    for i in range (1, n - 1):
        if not(s[i-1] not in vowels and s[i] in vowels and s[i+1] not in vowels):
            result.append(s[i])
            
    result.append(s[-1])

        
    return ''.join(result)
    