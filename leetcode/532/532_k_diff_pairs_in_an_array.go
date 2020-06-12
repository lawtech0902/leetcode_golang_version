/*
__author__ = 'lawtech'
__date__ = '2020/06/12 11:07 上午'
*/

package _532

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	numsHash := make(map[int]bool)
	diffHash := make(map[int]bool)

	for _, num := range nums {
		if numsHash[num-k] {
			diffHash[num-k] = true
		}

		if numsHash[num+k] {
			diffHash[num] = true
		}

		numsHash[num] = true
	}

	return len(diffHash)
}
