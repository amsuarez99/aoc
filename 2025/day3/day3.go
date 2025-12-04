package day3

import (
	"fmt"
	"math"
)

const zero = byte('0')
const newLn = byte('\n')

func findLargest(arr []byte) int {
	fmt.Printf("trying to find largest %s\n", string(arr))
	var largestIdx int
	for index, _ := range arr {
		if arr[index] > arr[largestIdx] {
			largestIdx = index
		}
	}
	return largestIdx
}

func add(largeLeft, largeRight byte) int {
	result := (int(largeLeft-zero) * 10) + int(largeRight-zero)
	return result
}

func Solution(input []byte) int {
	var bankLen, i, out, left, right int
	for input[i] != newLn {
		i += 1
	}

	bankLen = i
	// [1, 2, 4, 5]
	for right < len(input) {
		start := left
		right += 1
		i = right
		for i < start+bankLen-1 {
		}
		if input[right] > input[i] {
			out += add(input[left], input[right])
		} else {
			out += add(input[left], input[i])
		}
		// before \n
		right = i + 2
		left = right
	}

	return out
}

func Solution2(input []byte, k int) int {
	var bankLen, i, sum int
	for input[i] != newLn {
		i += 1
	}
	bankLen = i

	offset := 0
	for offset < len(input) {
		size := k
		fmt.Println("offset", offset)
		left := 0
		i = 0
		for i < bankLen {
			fmt.Println("i:", i)
			if i == bankLen-size {
				localLargestIdx := findLargest(input[offset+left : offset+i+1])
				largestIdx := left + localLargestIdx
				sum += int(input[offset+largestIdx]-zero) * int(math.Pow10(size-1))
				fmt.Println("Adding", left, i, sum)
				size -= 1
				i = largestIdx + 1
				left = i
			} else {
				i += 1
			}
		}
		offset += bankLen + 1
	}
	return sum
}
