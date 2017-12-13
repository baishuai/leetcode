package p299

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		a = b
	}
	return a
}

func getHint(secret string, guess string) string {

	if len(secret) != len(guess) {
		return "error"
	}
	a, b := 0, 0
	charsS := make(map[byte]int)
	charsG := make(map[byte]int)
	for i, c := range []byte(secret) {
		if c == guess[i] {
			a++
		} else {
			charsS[c]++
			charsG[guess[i]]++
		}
	}
	for k, v := range charsS {
		b = b + min(v, charsG[k])
	}
	return fmt.Sprintf("%dA%dB", a, b)
}
