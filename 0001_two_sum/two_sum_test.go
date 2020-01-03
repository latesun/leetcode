package two_sum

import "testing"

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, v := range nums {
		if _, ok := hashMap[target-v]; ok {
			return []int{i, hashMap[target-v]}
		}
		hashMap[v] = i
	}
	return []int{}
}

func BenchmarkTwoSum(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nums := []int{2, 7, 11, 15}
		target := 9
		twoSum(nums, target)
	}
}
