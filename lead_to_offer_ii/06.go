package lead_to_offer_ii

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int, len(numbers))
	for i, num := range numbers {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}

		m[num] = i
	}

	return nil
}
