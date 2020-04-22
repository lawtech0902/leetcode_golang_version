/*
__author__ = 'lawtech'
__date__ = '2020/04/22 2:27 下午'
*/

package _1081

func smallestSubsequence(text string) string {
	var (
		stack    []byte
		counts   [26]int
		inserted [26]bool
	)

	for i := 0; i < len(text); i++ {
		counts[text[i]-'a']++
	}

	for i := 0; i < len(text); i++ {
		counts[text[i]-'a']--
		if inserted[text[i]-'a'] {
			continue
		}

		for j := len(stack) - 1; j >= 0; j-- {
			if !(text[i] < stack[j] && counts[stack[j]-'a'] > 0) {
				break
			}

			inserted[stack[j]-'a'] = false
			stack = stack[:j]
		}

		inserted[text[i]-'a'] = true
		stack = append(stack, text[i])
	}

	return string(stack)
}
