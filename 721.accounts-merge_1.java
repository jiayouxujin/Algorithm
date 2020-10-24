/*
 * @lc app=leetcode id=721 lang=java
 *
 * [721] Accounts Merge
 */

// @lc code=start
class Solution {
    public List<List<String>> accountsMerge(List<List<String>> accounts) {
        int[] parent=new int[accounts.size()];
        for(int i=0;i<accounts.size();i++){
            parent[i]=i;
        }
        Map<String,Integer> owners=new HashMap<>();
        for(int i=0;i<accounts.size();i++){
            for(int j=1;j<accounts.get(i).size();j++){
                String email=accounts.get(i).get(j);
                if(owners.containsKey(email)){
                    int person=owners.get(email);
                    int p1=findParent(parent,i);
                    int p2=findParent(parent,person);
                    if(p1!=p2){
                        parent[p2]=p1;
                    }
                }else{
                    owners.put(email,i);
                }
            }
        }
        Map<Integer,TreeSet<String>> users=new HashMap<>();
        for(int i=0;i<accounts.size();i++){
            int p=findParent(parent,i);
            List<String> email=accounts.get(i);
            users.putIfAbsent(p,new TreeSet<String>());
            users.get(p).addAll(email.subList(1,email.size()));
        }

        List<List<String>> res=new ArrayList<>();
        for(Integer idx:users.keySet()){
            String name=accounts.get(idx).get(0);
            ArrayList<String> email=new ArrayList<>(users.get(idx));
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

