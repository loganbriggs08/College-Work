package main

import "Linear-Search/random"

var exampleArray *[]int

func main() {
	*exampleArray = random.RandomArray(1, 250, 10)
}
