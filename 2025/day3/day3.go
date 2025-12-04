package day3

import (
	"slices"
)

const zero = byte('0')
const newLn = byte('\n')

func findLargest(arr []byte) int {
	var largestIdx int
	for index := range arr {
		if arr[index] > arr[largestIdx] {
			largestIdx = index
		}
	}
	return largestIdx
}

func Solution(input []byte, k int) int {
	var sum int
	bankLen := slices.Index(input, newLn)

	offset := 0
	for offset < len(input) {
		localSum := 0
		end := offset + bankLen
		size := k
		leftBound := offset
		i := 0
		for i < end {
			if i == end-size {
				localLargestIdx := findLargest(input[leftBound : i+1])
				largestIdx := leftBound + localLargestIdx
				localSum = localSum*10 + int(input[largestIdx]-zero)
				size -= 1
				leftBound = largestIdx + 1
				i = leftBound
			} else {
				i += 1
			}
		}
		sum += localSum
		offset += bankLen + 1
	}
	return sum
}
