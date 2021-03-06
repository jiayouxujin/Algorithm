/*
 * @lc app=leetcode id=973 lang=golang
 *
 * [973] K Closest Points to Origin
 */

// @lc code=start
func kClosest(points [][]int, K int) [][]int {
    sort.Slice(points,func(i,j int)bool{
		return points[i][0]*points[i][0]+points[i][1]*points[i][1]<points[j][0]*points[j][0]+points[j][1]*points[j][1]
	})
	ans:=make([][]int,K)
	for i:=0;i<K;i++{
		ans[i]=points[i]
	}
	return ans
}
// @lc code=end

