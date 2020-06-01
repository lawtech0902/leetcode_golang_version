/*
__author__ = 'lawtech'
__date__ = '2020/06/01 10:32 上午'
*/

package _491

func findSubsequences(nums []int) [][]int {
	var (
		paths [][]int
		path  []int
		dfs   func(loc int, length int, preNum int)
	)

	dfs = func(loc int, length int, preNum int) {
		if length >= 2 {
			ll := len(path)
			tmp := make([]int, ll)
			copy(tmp, path)
			paths = append(paths, tmp)
		}

		// 以一个for循环为一层，剪枝后面出现的同层元素
		mmap := make(map[int]int) // 记录当前层已经被访问过的值，后面再次遇到则此值必然重复，因为在前面更深的递归调用中已经写入了

		for i := loc; i < len(nums); i++ {
			_, ok := mmap[nums[i]]
			if nums[i] < preNum || ok {
				continue
			}
			mmap[nums[i]] = 1
			path = append(path, nums[i])
			dfs(i+1, length+1, nums[i])
			path = path[:len(path)-1]
		}
	}

	dfs(0, 0, -101)

	return paths
}
