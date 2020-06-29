/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-01-03 16:19:42
 */

package _605

func canPlaceFlowers(flowerbed []int, n int) bool {
	size := len(flowerbed)

	for i := 0; i < size; i++ {
		if flowerbed[i] == 0 &&
			((i+1 < size && flowerbed[i+1] == 0) || i+1 >= size) &&
			((i-1 >= 0 && flowerbed[i-1] == 0) || i-1 < 0) {
			flowerbed[i] = 1
			n--
			if n <= 0 {
				return true
			}
		}
	}

	return n <= 0
}
