package main

import (
	"fmt"
	"time"
)

// Sort ascending input array
func bubbleSort(c []int) {
	if c == nil {
		return
	}

	if len(c) == 1 {
		return
	}

	for i := 0; i < len(c)-1; i++ {
		for j := 0; j < len(c)-i-1; j++ {
			if c[j] > c[j+1] {
				//swap elements
				temp := c[j+1]
				c[j+1] = c[j]
				c[j] = temp
			}
			fmt.Printf("\r")
			fmt.Printf("Sorting container = %v", c)
			time.Sleep(100 * time.Millisecond)
		}
	}
	fmt.Printf("\r") //temporary not nice solution
}

func main() {
	fmt.Println("Bubble sort")
	collection := []int{13, 4, 35, 104, 28, 6, 12, 200, 1, -5, 44, 55, 11, 99, -9, 21, -100, 67, -37, 12}
	fmt.Printf("Initial container = %v\n", collection)
	bubbleSort(collection)
	fmt.Printf("Sorted container  = %v\n", collection)
}
