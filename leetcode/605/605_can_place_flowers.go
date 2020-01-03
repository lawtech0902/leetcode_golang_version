package _605

/*
__author__ = 'lawtech'
__date__ = '2018/8/20 下午11:05'
*/

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
