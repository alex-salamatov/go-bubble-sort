package main

import (
	"math/rand/v2"
)

func makeIntArray(maxVal int, maxLen int) []int {
	arrLen := rand.IntN(maxLen)
	arr := make([]int, arrLen)

	for i := 0; i < arrLen; i++ {
		arr[i] = rand.IntN(maxVal)
		if rand.IntN(2) == 1 {
			arr[i] = -1 * arr[i]
		}
	}

	return arr
}
