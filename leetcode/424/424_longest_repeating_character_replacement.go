/*
__author__ = 'lawtech'
__date__ = '2020/05/15 4:47 下午'
*/

package _424

func characterReplacement(s string, k int) int {
	var maxLength int
	for i := 0; i < 26; i++ {
		ops := k
		start, end := 0, 0
		for end = 0; end < len(s); end++ {
			if s[end]-'A' != uint8(i) {
				for ops == 0 {
					if s[start]-'A' != uint8(i) {
						ops++
					}

					start++
				}
				ops--
			}

			if end-start+1 > maxLength {
				maxLength = end - start + 1
			}
		}
	}

	return maxLength
}
