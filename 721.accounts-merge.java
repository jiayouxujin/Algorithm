/*
 * @lc app=leetcode id=721 lang=java
 *
 * [721] Accounts Merge
 */

// @lc code=start
class Solution {
    public List<List<String>> accountsMerge(List<List<String>> accounts) {
        int[] parents=new int[accounts.size()];
        for(int i=0;i<accounts.size();i++){
            parents[i]=i;
        }

        Map<String,Integer> owner=new HashMap<>();
        for(int i=0;i<accounts.size();i++){
            for(int j=1;j<accounts.get(i).size();j++){
                String email=accounts.get(i).get(j);
                if(owner.containsKey(email)){
                    int p=owner.get(email);
                    int p1=findParent(parents,p);
                    int p2=findParent(parents,i);
                    if(p2!=p1){
                        parents[p1]=p2;
                    }
                }else{
                    owner.put(email,i);
                }
            }
        }

        Map<Integer,TreeSet<String>> map=new HashMap<>();
        for(int i=0;i<accounts.size();i++){
            int p=findParent(parents,i);
            List<String> email=accounts.get(i);
            map.putIfAbsent(p,new TreeSet<String>());
            map.get(p).addAll(email.subList(1,email.size()));
        }

        List<List<String>> res=new ArrayList<>();
        for(Integer idx:map.keySet()){
            String name=accounts.get(idx).get(0);
            ArrayList<String> email=new ArrayList(map.get(idx));
            email.add(0,name);
            res.add(email);
        }
        return res;
    }

    private int findParent(int[] parents,int idx){
        while(idx!=parents[idx]){
            parents[idx]=parents[parents[idx]];
            idx=parents[idx];
        }
        return idx;
    }
}
// @lc code=end

