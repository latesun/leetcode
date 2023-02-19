package twosum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for curIdx, num := range nums {
		if firstIdx, ok := m[target-num]; ok {
			return []int{firstIdx, curIdx}
		}
		m[num] = curIdx
	}
	return nil
}
