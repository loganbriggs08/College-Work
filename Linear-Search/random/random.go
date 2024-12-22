package random

import (
	"math/rand"
)

func randomNumber(minValue int, maxValue int) int {
	return rand.Intn(maxValue-minValue) + minValue
}

func RandomArray(minValue int, maxValue int, length int) []int {
	randomNumbers := make([]int, length)

	for currentIndex := 0; currentIndex <= length; currentIndex++ {
		randomNumbers[currentIndex] = randomNumber(minValue, maxValue)
	}

	return randomNumbers
}
