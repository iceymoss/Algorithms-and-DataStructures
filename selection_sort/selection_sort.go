package selectionsort

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		tag := i
		for j := i + 1; j < len(arr); j++ {
			if arr[tag] > arr[j] {
				tag = j
			}
		}
		arr[i], arr[tag] = arr[tag], arr[i]
	}
	return arr
}
