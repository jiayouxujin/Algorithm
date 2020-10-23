/*
 * @lc app=leetcode id=547 lang=java
 *
 * [547] Friend Circles
 */

// @lc code=start
class Solution {
    public int findCircleNum(int[][] M) {
        if(M==null||M.length==0){
            return 0;
        }

        int ans=0;
        boolean[] visit=new boolean[M.length];
        for(int i=0;i<M.length;i++){
            if(!visit[i]){
                dfs(M,visit,i);
                ans++;
            }
        }
        return ans;
    }

    public void dfs(int[][] M,boolean[] v,int index){
        v[index]=true;
        for(int i=0;i<M[0].length;i++){
            if(M[index][i]==1&&!v[i]){
                dfs(M,v,i);
            }
        }
    }
}
// @lc code=end

