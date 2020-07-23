/*
__author__ = 'lawtech'
__date__ = '2020/07/23 11:14 上午'
*/

package _771

func numJewelsInStones(J, S string) int {
	if len(S) == 0 || len(J) == 0 {
		return 0
	}

	jMap := make(map[rune]bool, len(J))
	for _, j := range J {
		jMap[j] = true
	}

	res := 0
	for _, s := range S {
		if jMap[s] {
			res++
		}
	}

	return res
}
