class Solution:
    def twoSum(self, nums: list, target: int) -> list:
        hash_map = {}
        for i, v in enumerate(nums):
            if target - v in hash_map.keys():
                return [i, hash_map[target - v]]
            hash_map[v] = i


if __name__ == '__main__':
    nums = [2, 7, 11, 15]
    target = 9
    s = Solution()
    print(s.twoSum(nums, target))
