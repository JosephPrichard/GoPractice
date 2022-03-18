package sort

func BubbleSort(array []int) {
	cont := true
	for cont {
		cont = false
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
				cont = true
			}
		}
	}
}

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, left int, right int) {
	if right <= left {
		return
	}

	mid := left + (right-left)/2

	mergeSort(arr, left, mid)
	mergeSort(arr, mid+1, right)
	SortedMerge(arr, left, mid, right)
}

func SortedMerge(arr []int, left int, mid int, right int) {
	mergedArr := make([]int, right-left+1)

	l := left
	r := mid + 1
	m := 0

	for l < mid+1 && r < right+1 {
		if arr[l] < arr[r] {
			mergedArr[m] = arr[l]
			l++
		} else {
			mergedArr[m] = arr[r]
			r++
		}
		m++
	}

	for l < mid+1 {
		mergedArr[m] = arr[l]
		l++
		m++
	}

	for r < right+1 {
		mergedArr[m] = arr[r]
		r++
		m++
	}

	m = left
	for i := 0; i < len(mergedArr); i++ {
		arr[m] = mergedArr[i]
		m++
	}
}
