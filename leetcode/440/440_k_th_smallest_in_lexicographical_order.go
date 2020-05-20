/*
__author__ = 'lawtech'
__date__ = '2020/05/20 2:57 下午'
*/

package _440

func findKthNumber(n int, k int) int {
	currentNode := 1
	prefix := 1

	for currentNode < k {
		count := getCount(prefix, n)
		if currentNode+count > k {
			prefix *= 10
			currentNode++
		} else {
			prefix++
			currentNode += count
		}
	}

	return prefix
}

func getCount(prefix, n int) int {
	count := 0
	cur := prefix
	next := prefix + 1

	for cur <= n {
		count += min(next, n+1) - cur
		cur *= 10
		next *= 10
	}

	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
