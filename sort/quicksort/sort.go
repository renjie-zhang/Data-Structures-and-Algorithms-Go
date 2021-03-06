package quicksort

import (
	"reflect"
	"testing"
)

func QuickSort(arr []int, low int, high int) {
	if low > high {
		return
	}
	j := part(arr, low, high)
	QuickSort(arr, low, j-1)
	QuickSort(arr, j+1, high)
}

func part(arr []int, low int, high int) int {
	i := low
	// 使用最右边作为基准
	pivot := arr[high]
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func Equal(t *testing.T, expected, result interface{}) {
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("should be %v instead of %v", expected, result)
	}
}
