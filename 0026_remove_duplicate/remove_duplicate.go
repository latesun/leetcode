package removeduplicate

func removeDuplicates(nums []int) int {
	cnt := 0
	for i, num := range nums {
		if i == 0 || nums[i] != nums[i-1] {
			nums[cnt] = num
			cnt++
		}
	}

	return cnt
}
