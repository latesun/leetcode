package searchmatrix

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	n, left, right := len(matrix[0]), 0, len(matrix)*len(matrix[0])-1
	for left <= right {
		mid := left + (right-left)/2
		x := matrix[mid/n][mid%n]

		if x == target {
			return true
		}

		if x < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
