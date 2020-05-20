/*
__author__ = 'lawtech'
__date__ = '2020/05/20 10:01 上午'
*/

package _438

// 滑动窗口
func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(s) < len(p) {
		return res
	}

	hash := [26]int{}
	for i := range p {
		hash[p[i]-'a'] += 1
	}

	left, right := 0, 0
	count := len(p)

	for right < len(s) {
		if hash[s[right]-'a'] > 0 {
			count--
		}

		hash[s[right]-'a']--
		right++

		if count == 0 {
			res = append(res, left)
		}

		if right-left == len(p) {
			if hash[s[left]-'a'] >= 0 {
				count++
			}

			hash[s[left]-'a']++
			left++
		}
	}

	return res
}
