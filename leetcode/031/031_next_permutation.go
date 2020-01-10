package _31

/*
__author__ = 'lawtech'
__date__ = '2018/8/18 下午9:06'
*/

/*
1　　2　　7　　4　　3　　1
下一个排列为：
1　　3　　1　　2　　4　　7
那么是如何得到的呢，步骤如下：
从后往前，找到最长的降序排列
1　　2　　7　　4　　3　　1
把这个降序排列，转换成升序排列
1　　2　　1　　3　　4　　7
把序列前的元素，与序列中，第一个大于他的元素互换。
1　　3　　1　　2　　4　　7
*/

func nextPermutation(a []int) {
	// 找到递减序列a[left+1:]
	left := len(a) - 2
	for 0 <= left && a[left] >= a[left+1] {
		left--
	}

	// 逆转a[left+1:]
	reverse(a, left+1)

	if left == -1 {
		return
	}

	right := search(a, left+1, a[left])
	a[left], a[right] = a[right], a[left]
}

func reverse(a []int, l int) {
	r := len(a) - 1

	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}

func search(a []int, l, target int) int {
	r := len(a) - 1
	l--

	for l+1 < r {
		mid := (l + r) / 2
		if target < a[mid] {
			r = mid
		} else {
			l = mid
		}
	}

	return r
}
