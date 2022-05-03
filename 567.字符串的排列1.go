/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
    if s1==""{
        return true
    }else if s2==""{
        return false
    }

    need,window:=make(map[byte]int),make(map[byte]int)
    for i:=0;i<len(s1);i++{
        need[s1[i]]++
    }

    left,right,count:=0,0,0
    for right<len(s2){
        c:=s2[right]
        right++

        if v,ok:=need[c];ok{
            window[c]++
            if window[c]==v{
                count++
            }
        }

        for right-left>=len(s1){
            if count==len(need){
                return true
            }

            d:=s2[left]
            left++

            if v,ok:=need[d];ok{
                if window[d]==v{
                    count--
                }
                window[d]--
            }
        }
    }
    return false
}
// @lc code=end

