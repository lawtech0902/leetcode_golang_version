/*
__author__ = "lawtech"
__date__ = "2020/1/9 5:33 下午"
*/

package _17

func letterCombinations(digits string) []string {
	digitMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	size := len(digits)
	if size == 0 {
		return []string{}
	}

	res := []string{""}
	for _, char := range digits {
		var tmp []string
		for _, y := range res {
			for _, x := range digitMap[string(char)] {
				tmp = append(tmp, y+string(x))
			}
			res = tmp
		}
	}

	return res
}
