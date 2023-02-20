package sort

func BubbleSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := 0; i < len(arr); i++ {
		swap := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swap = true
			}
		}

		if !swap {
			break
		}
	}
}
