package main

import (
	"fmt"
)

func set(slice []int) []int {
	valueMap := make(map[int]bool)
	results := []int{}
	for i := 0; i < len(slice); i++ {
		if _, found := valueMap[slice[i]]; !found {
			valueMap[slice[i]] = true
			results = append(results, slice[i])
		}
	}
	return results
}

func reverseSlice(slice []int) []int {
	count := len(slice) - 1
	newSlice := []int{}
	for i := count; i >= 0; i-- {
		newSlice = append(newSlice, slice[i])
	}
	return newSlice
}

func main() {
	sampleList := []int{10, 20, 30, 5, 9, 10, 40, 40, 50, 40, 20, 45, 50, 20}
	// expectedResult := []int{20 50 45 40 10}
	_m := make(map[int]int)
	count := len(sampleList) - 1
	// Number of results
	numberToRetrieve := 5
	resultList := []int{}
	// This flow ensures that we only iterate over a subset of the slice in reverse
	for i := count; i >= 0; i-- {
		// Check if we still need more values to populate our set
		if numberToRetrieve != 0 {
			// Check our map to see if the current value specified in the sampleList[i] exists in the map
			if _, found := _m[sampleList[i]]; !found {
				// Assign to our map if not found, also, we can save true as the value instead of the int
				_m[sampleList[i]] = sampleList[i]
				// We need to append the sampleList[i] value to our resultList, creating a unique set.
				resultList = append(resultList, sampleList[i])
				// Because we added a value to our set, we can decrement the required numberToRetrieve int
				numberToRetrieve--
			}
		} else {
			// Exit this loop once we have our numberToRetrieve zeroed out.
			break
		}
	}

	fmt.Println("Slice ", resultList)

	// Cleaner implementation
	sliceFromFunc := []int{}
	// Reverse the sampleList so that we can loop on it without having to worry about the reverse ordering
	// when having to create our set.
	reversedSlice := reverseSlice(sampleList)
	/* Ideally we would provide a subset of a slice slice instead of an entire list of n values.
	   I would imagine whatever calls on this method would provide this generated list of values.
	   The below code would also provide a subset slice of the original array. */
	// reversedSlice := reverseSlice(sampleList[len(sampleList)-(1+5):])
	// Once we have our reversed slice, we can now create a set with unique values
	sliceFromFunc = set(reversedSlice)
	fmt.Println("result slice", sliceFromFunc[:5])
}
