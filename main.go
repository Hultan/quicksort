package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	var numbers []int
	var numbers2 []int

	// Create a slice of random numbers
	for i := 0; i < 100000; i++ {
		n := rand.Intn(10000)
		numbers = append(numbers, n)
	}

	// Print unsorted array
	// fmt.Println(numbers)

	// Create a copy of the slice so that we can compare it with
	// Go:s internal sort.Ints()-method.
	numbers2 = make([]int, len(numbers))
	copy(numbers2, numbers)

	// Sort slice using my method
	time1 := time.Now()
	quicksort(numbers, 0, len(numbers)-1)
	fmt.Println("(MY SORT) Sorted in ", time.Now().Sub(time1).Milliseconds(), "ms!")
	// fmt.Println(numbers)

	// Sort slice using Go:s sort.Ints() method
	time2 := time.Now()
	sort.Ints(numbers2)
	fmt.Println("(GOLANG) Sorted in ", time.Now().Sub(time2).Milliseconds(), "ms!")
	// fmt.Println(numbers2)

	// Just make sure it is sorted correctly
	for i := range numbers {
		if numbers[i] != numbers2[i] {
			fmt.Println("FAILURE!!!")
		}
	}
}

func quicksort(numbers []int, lowIndex, highIndex int) {
	// Stop recursion when size of array is 1 (or less)
	if lowIndex >= highIndex {
		return
	}

	// Choose a pivot point
	// For simplicity, just pick the last number
	pivot := numbers[highIndex]

	// Left and right pointers
	leftPointer := lowIndex
	rightPointer := highIndex

	for leftPointer < rightPointer {
		// Step leftPointer to the right, until we
		// find a number larger or equal to the pivot
		for numbers[leftPointer] < pivot && leftPointer < rightPointer {
			leftPointer++
		}

		// Step rightPointer to the left, until we
		// find a number smaller than the pivot
		for numbers[rightPointer] >= pivot && leftPointer < rightPointer {
			rightPointer--
		}

		// Swap leftPointer and rightPointer numbers
		numbers[leftPointer], numbers[rightPointer] = numbers[rightPointer], numbers[leftPointer]
	}

	// Swap pivot with leftPointer
	numbers[leftPointer], numbers[highIndex] = numbers[highIndex], numbers[leftPointer]

	// Recursively sort left and right side
	quicksort(numbers, lowIndex, leftPointer-1)
	quicksort(numbers, leftPointer+1, highIndex)
}
