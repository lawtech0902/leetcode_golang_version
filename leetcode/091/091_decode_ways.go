/*
__author__ = 'lawtech'
__date__ = '2020/1/17 5:06 下午'
*/

package _91

import "strconv"

func numDecodings(s string) int {
	/*
		当s[i-2:i]这两个字符是10~26但不包括10和20这两个数时，比如21，那么可以有两种编码方式（BA，U），所以dp[i]=dp[i-1]+dp[i-2]
		当s[i-2:i]等于10或者20时，由于10和20只有一种编码方式，所以dp[i]=dp[i-2]
		当s[i-2:i]不在以上两个范围时，如09这种，编码方式为0，而31这种，dp[i]=dp[i-1]。
		注意初始化时：dp[0]=1,dp[1]=1
	*/
	if s == "" || s[0] == '0' {
		return 0
	}

	dp := []int{1, 1}
	for i := 2; i < len(s)+1; i++ {
		num, _ := strconv.Atoi(s[i-2 : i])
		if 10 <= num && num <= 26 && s[i-1] != '0' {
			dp = append(dp, dp[i-1]+dp[i-2])
		} else if num == 10 || num == 20 {
			dp = append(dp, dp[i-2])
		} else if s[i-1] != '0' {
			dp = append(dp, dp[i-1])
		} else {
			return 0
		}
	}

	return dp[len(s)]
}
