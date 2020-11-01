/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {
	var res []string
	
	gen(&res,n,n,"");
	return res;
}
func gen(res *[]string,left,right int,tmp string)  {
	if left==0&&right==0{
		*res=append(*res,tmp)
		return ;
	}
	if left>0{
		gen(res,left-1,right,tmp+"(")
	}

	if right>left{
		gen(res,left,right-1,tmp+")")
	}
}
// @lc code=end

