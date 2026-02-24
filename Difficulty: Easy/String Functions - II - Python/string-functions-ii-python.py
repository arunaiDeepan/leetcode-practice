# Function to check if string
# starts and ends with 'gfg'
def gfg(S):
    s_lower = S.lower()
    if (s_lower.startswith("gfg") and s_lower.endswith("gfg")):
        print("Yes")
    else:
        print("No")