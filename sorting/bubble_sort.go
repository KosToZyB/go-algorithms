package main

import (
	"fmt"
	"math/rand"
	"time"
)

//SortBubble function to implement bubble sort
func SortBubble(data []int) {
	n := len(data)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func main() {
	arraySize := 20
	maxGenerateNumber := 10

	data := make([]int, arraySize)

	rand.Seed(time.Now().Unix())
	for d := range data {
		data[d] = rand.Intn(maxGenerateNumber)
	}

	fmt.Println("Print source array: ", data)
	SortBubble(data)
	fmt.Println("Print after sorting: ", data)
}
