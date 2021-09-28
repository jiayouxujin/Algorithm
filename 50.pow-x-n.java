/*
 * @lc app=leetcode.cn id=50 lang=java
 *
 * [50] Pow(x, n)
 */

// @lc code=start
class Solution {
    public double myPow(double x, int n) {
        long N = n;
        if (N < 0) {
            N = -N;
            x = 1 / x;
        }
        double ans = 1;
        double current = x;
        while (N > 0) {
            if ((N & 1) == 1) {
                ans *= current;
            }
            current = current * current;
            N /= 2;
        }
        return ans;
    }
}
// @lc code=end
