package _06

/*
__author__ = 'lawtech'
__date__ = '2020/1/6 4:40 下午'
*/

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var (
		cs    []uint8
		count int
		index int
	)

	for i := 0; i < numRows; i++ {
		count = 0
		index = count*(2*numRows-2) + i
		for index < len(s) {
			cs = append(cs, s[index])
			if i != 0 && i != numRows-1 {
				index = (count+1)*(2*numRows-2) - i
				if index >= len(s) {
					break
				}
				cs = append(cs, s[index])
			}
			count++
			index = count*(2*numRows-2) + i
		}
	}

	return string(cs)
}
