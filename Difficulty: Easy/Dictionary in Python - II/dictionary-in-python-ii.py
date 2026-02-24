def pair_sum(arr, sum):
    seen = {}
    
    for i in arr:
        
        c = sum - i
        
        if c in seen:
            return True
        
        seen[i] = True
    
    return False