// Galang - Golang common utilities
// Copyright (c) 2020-present, gakkiiyomi@gamil.com
//
// gakkiyomi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package sort

import (
	"github.com/gakkiyomi/galang/structure"
)

//SelectionSort is the realization of selection sort
func SelectionSort(source []int) {
	length := len(source)
	for i := 0; i < length; i++ {
		maxIndex := 0
		for j := 1; j < length-i; j++ {
			if source[j] > source[maxIndex] {
				maxIndex = j
			}
		}
		source[length-i-1], source[maxIndex] = source[maxIndex], source[length-i-1]
	}

}

//SelectionSort is the realization of bubble sort
func BubbleSort(source []int) {

	for i := 0; i < len(source); i++ {
		for j := 1; j < len(source)-i; j++ {
			if source[j] < source[j-1] {
				source[j], source[j-1] = source[j-1], source[j]
			}
		}
	}

}

//SelectionSort is the realization of insertion sort
func InsertionSort(source []int) {
	length := len(source)

	for i := 1; i < length; i++ {
		j := i
		for j > 0 {
			if source[j-1] > source[j] {
				source[j], source[j-1] = source[j-1], source[j]
			}
			j -= 1
		}
	}
}

//SelectionSort is the realization of quick sort
func QuickSort(source []int) {
	recursionSort(source, 0, len(source)-1)
}

func recursionSort(source []int, left, right int) {
	if left < right {
		index := partition(source, left, right)
		recursionSort(source, left, index-1)
		recursionSort(source, index+1, right)
	}
}

func partition(source []int, left, right int) int {
	for left < right {
		for left < right && source[left] <= source[right] {
			right--
		}
		if left < right {
			source[left], source[right] = source[right], source[left]
			left++
		}

		for left < right && source[left] <= source[right] {
			left++
		}
		if left < right {
			source[left], source[right] = source[right], source[left]
			right--
		}
	}

	return left //or right
}

func HeapSort(source []int, asc bool) []int {

	result := []int{}

	if asc {

		minHeap := structure.NewMinHeap(source)
		for range source {
			// 也可以使用栈或者队列
			//result = array.Array.InsertAtIndexByIntArray(result, maxHeap.ExtractMax(), 0)
			result = append(result, minHeap.ExtractMin())
		}

	} else {
		maxHeap := structure.NewMaxHeap(source)
		for range source {
			result = append(result, maxHeap.ExtractMax())
		}

	}
	return result
}
