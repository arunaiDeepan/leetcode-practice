def friends_in_trouble(j_angry, s_angry):
    return (j_angry & s_angry) or (not (j_angry ^ s_angry))