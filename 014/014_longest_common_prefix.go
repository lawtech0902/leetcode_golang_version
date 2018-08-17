package _14

/*
__author__ = 'lawtech'
__date__ = '2018/8/18 上午12:52'
*/

func longestCommonPrefix(strs []string) string {
	short := shortestStr(strs)

	for i, r := range short {
		for j := 0; j < len(strs); j ++ {
			if strs[j][i] != byte(r) {
				return strs[j][:i]
			}
		}
	}

	return short
}

func shortestStr(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]
	for _, s := range strs {
		if len(res) > len(s) {
			res = s
		}
	}

	return res
}
