class Solution:
    def sort012(self, arr):
        r = len(arr) - 1
        l = 0
        m = 0
        while m <= r:
            if arr[m] == 0:
                arr[l], arr[m] = arr[m], arr[l]
                m += 1
                l += 1
            elif arr[m] == 1:
                m += 1
            else:
                arr[m], arr[r] = arr[r], arr[m]
                r -= 1
                
        return arr