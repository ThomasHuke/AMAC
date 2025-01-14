// 快排 归并 希尔排序
package main

func XiEr(x ...int) []int {
	var r []int
	for i := 0; i < len(x); i = i + 10 {
		r = Cha(x[0:i]...)
		if i < len(x) && len(x)-i < 10 {
			r = Cha(x...)
		}
	}

	return r
}

func ShellSort(arr []int) []int {
	length := len(arr)
	gap := 1
	for gap < gap/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 3
	}
	return arr
}

// 归并排序

func MergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}

	left := arr[0 : n/2]
	right := arr[n/2:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}

// 快排
func QuickSort(x []int) []int {
	left := 0
	right := len(x) - 1
	return quickSort(left, right, x)
}

func quickSort(left int, right int, arr []int) []int {
	if left < right {
		m := pattion(left, right, arr)
		quickSort(left, m-1, arr)
		quickSort(m+1, right, arr)
	}
	return arr
}
func pattion(left int, right int, arr []int) int {
	p := left
	index := p + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[p] {
			arr[i], arr[index] = arr[index], arr[i]
			index += 1
		}
	}
	arr[p], arr[index-1] = arr[index-1], arr[p]
	return index - 1
}
