class Solution {
public:
    vector<int> getRow(int rowIndex) {
        int n = rowIndex;
        vector<vector<int>> res(n+1);
        
        for(int i=0;i<=n;i++){
            res[i].resize(i+1);
            res[i][0]=res[i][i]=1;
            for(int j=1;j<i;j++)
                res[i][j] = (res[i-1][j-1]+res[i-1][j]);
                
            
        }
        vector<int> ans;
        for(int i=0;i<=n;i++){
            ans.push_back(res[n][i]);
        }
        return ans;
    }
};