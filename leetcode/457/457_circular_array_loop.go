/*
__author__ = 'lawtech'
__date__ = '2020/05/25 9:58 上午'
*/

package _457

// nextIdx returns the next index in nums to move from curIdx.
func nextIdx(nums []int, curIdx int) int {
	next := curIdx + nums[curIdx]
	next %= len(nums)
	if next < 0 {
		return len(nums) + next
	}

	return next
}

func circularArrayLoop(nums []int) bool {
	n := len(nums)
	// Array of size 1 couldn't have qualified cycle because the cycle length is 1.
	// This is required to maks "next %= len(nums)" work. Consider the test case "[-1]"
	// and see what nextIdx(nums, 0) will return.
	if n == 1 {
		return false
	}

	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		// Find first non-visited element.
		if visited[i] {
			continue
		}

		visited[i] = true

		// Run Floyd's algorithm.
		//
		// See https://www.youtube.com/watch?v=zbozWoMgKW0.
		hare := i
		tortoise := i
		for {
			// Hare moves two steps.
			hare = nextIdx(nums, hare)
			visited[hare] = true
			hare = nextIdx(nums, hare)
			visited[hare] = true

			// Tortoise moves one step.
			tortoise = nextIdx(nums, tortoise)
			visited[tortoise] = true

			// Because the array is circular (as promised by question), there's always a cycle.
			if hare == tortoise {
				break
			}
		}

		// Now find the start of cycle.
		//
		// See (the awesome) https://www.youtube.com/watch?v=LUm2ABqAs1w to know why this works.
		loopStart := i
		for loopStart != tortoise {
			loopStart = nextIdx(nums, loopStart)
			tortoise = nextIdx(nums, tortoise)
		}

		// Now check the cycle.
		cycle := nextIdx(nums, loopStart)
		if cycle == loopStart {
			// Cycle's length == 1
			continue
		}

		forwardCycle := nums[loopStart] > 0
		for cycle != loopStart {
			if (nums[cycle] > 0) != forwardCycle {
				break
			}
			cycle = nextIdx(nums, cycle)
		}

		// If we finish walking the cycle, the movements in the cycle are all either
		// forward or backward.
		if cycle == loopStart {
			return true
		}
	}

	return false
}
