/*
 * @Description:
 * @Author: lawtech
 * @Date: 2020-07-02 11:07:01
 */

package _638

// memo + dfs
func shoppingOffers(price []int, special [][]int, needs []int) int {
	maps := make(map[string]int)
	return shopping(price, special, needs, maps)
}

func shopping(price []int, special [][]int, needs []int, maps map[string]int) int {
	s := sliceToString(needs)
	if val, ok := maps[s]; ok {
		return val
	}

	res := 0
	for k := 0; k < len(needs); k++ {
		res += needs[k] * price[k]
	}

	i := 0
	for _, spec := range special {
		clone := make([]int, len(needs))
		for i = 0; i < len(needs); i++ {
			if spec[i] > needs[i] {
				break
			}
			clone[i] = needs[i] - spec[i]
		}

		if i == len(needs) {
			res = Min(res, spec[i]+shopping(price, special, clone, maps))
		}
	}
	maps[s] = res
	return res
}

func sliceToString(a []int) string {
	s := ""
	for _, v := range a {
		s += string(v) + ","
	}
	return s
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
