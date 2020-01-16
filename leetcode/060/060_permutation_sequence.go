/*
__author__ = 'lawtech'
__date__ = '2020/1/16 1:35 下午'
*/

package _60

func getPermutation(n int, k int) string {
	nums := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	s := nums[:n]
	r := ""

	for i := 0; i < n; i++ {
		selected := S(len(s), k) % len(s)
		r = r + nums[selected]
		s = append(s[0:selected], s[selected+1:]...)

		if len(s) == 1 {
			return r + s[0]
		}
	}

	return r
}

func F(n int) int {
	r := n

	for n > 1 {
		n--
		r = r * n
	}

	return r
}

func S(n, k int) int {
	// prevent divisor from 0
	if n-1 == 0 {
		return k - 1
	}

	// k needs to be subtracted 1 because the starting number is 1 and not 0
	return (k - 1) / F(n-1)
}
