def maxlength(s):
    max_l_1 = 0
    count = 0
    
    for i in s:
        if i == "1" :
            count += 1
        else:
            max_l_1 = max(max_l_1, count)
            count = 0
            
    return max(count, max_l_1)