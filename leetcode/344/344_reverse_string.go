package _344

/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午10:00'
*/

func reverseString(s string) string {
	bytes := []byte(s)

	i, j := 0, len(s)-1

	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i ++
		j --
	}

	return string(bytes)
}
