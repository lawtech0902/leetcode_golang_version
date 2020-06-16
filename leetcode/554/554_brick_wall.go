/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-06-16 10:33:57
 */

package _554

func leastBricks(wall [][]int) int {
	hash := make(map[int]int)
	maxLen := 0

	for i := 0; i < len(wall); i++ {
		sum := 0
		for j := 0; j < len(wall[i])-1; j++ {
			sum += wall[i][j]
			hash[sum] += 1
			if hash[sum] > maxLen {
				maxLen = hash[sum]
			}
		}
	}

	return len(wall) - maxLen
}
