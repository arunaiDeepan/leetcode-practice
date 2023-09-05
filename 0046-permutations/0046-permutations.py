class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        res = []
        self.bTrack(nums, 0, res)
        return res
    def bTrack(self, nums, s, res):
        if s == len(nums):
            res.append(nums[:])
            return

        for i in range(s, len(nums)):
            nums[s], nums[i] = nums[i], nums[s]
            self.bTrack(nums, s + 1, res)
            nums[s], nums[i] = nums[i], nums[s]