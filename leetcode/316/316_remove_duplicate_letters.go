/*
__author__ = 'lawtech'
__date__ = '2020/04/22 1:54 下午'
*/

package _316

// 贪心 + 栈 O(n) O(1)
func removeDuplicateLetters(s string) string {
	var (
		stack    []byte
		counts   [26]int
		inserted [26]bool
	)

	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']--
		if inserted[s[i]-'a'] {
			continue
		}

		for j := len(stack) - 1; j >= 0; j-- {
			if !(s[i] < stack[j] && counts[stack[j]-'a'] > 0) {
				break
			}

			inserted[stack[j]-'a'] = false
			stack = stack[:j]
		}

		inserted[s[i]-'a'] = true
		stack = append(stack, s[i])
	}

	return string(stack)
}
