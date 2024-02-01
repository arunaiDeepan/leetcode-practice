class Solution {
public:
    vector<vector<int>> divideArray(vector<int>& nums, int k) {
        vector<vector<int>> solution;
        sort(nums.begin(), nums.end());
        for(int i=2 ; i<nums.size(); i +=3){
            if(nums[i] - nums[i-2] <= k){
                solution.push_back({nums[i], nums[i-1], nums[i-2]});
            }
            else
                return {};
        }
        return solution;
    }
};