package main

import "fmt"

// 快速排序

func QuickSort(left int, right int, array *[6]int) {

	l := left
	r := right
	pivot := array[(left+right)>>2]
	temp := 0

	for l < r {
		for array[l] < pivot {
			l++
		}
		for array[l] > pivot {
			r--
		}
		if l >= r {
			break

		}
		temp = array[l]
		array[l] = array[r]
		array[r] = temp

		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		QuickSort(left, r, array)
	}
	if right > l {
		QuickSort(l, right, array)
	}

}
func main() {

	arr := [6]int{-9, 78, 0, 34, -567, 79}
	fmt.Println(arr)
	QuickSort(0, len(arr)-1, &arr)
	fmt.Println(arr)

}
