package twoSum

func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	targetMap := make(map[int]int, len(nums))
	for index, k := range nums {
		if v, ok := targetMap[k]; ok {
			return []int{v, index}
		}
		targetMap[target-k] = index
	}
	return res
}
