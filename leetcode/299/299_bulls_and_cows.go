/*
__author__ = 'lawtech'
__date__ = '2020/04/20 2:41 下午'
*/

package _299

import "fmt"

func getHint(secret string, guess string) string {
	cowHash := map[rune]int{}
	cowCount := 0
	bullCount := 0

	secretNums := []rune(secret)
	guessNums := []rune(guess)
	for i := 0; i < len(secret); i++ {
		if secretNums[i] != guessNums[i] {
			cowHash[secretNums[i]]++
		}
	}

	for i := 0; i < len(secret); i++ {
		if secretNums[i] == guessNums[i] {
			bullCount++
			continue
		}

		if v, ok := cowHash[guessNums[i]]; ok && v > 0 {
			cowHash[guessNums[i]]--
			cowCount++
		}
	}

	return fmt.Sprintf("%dA%dB", bullCount, cowCount)
}
