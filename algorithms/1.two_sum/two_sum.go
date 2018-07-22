package twosum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, n := range nums {
		j, ok := m[n]
		if ok {
			return []int{j, i}
		}
		m[target-nums[i]] = i
	}
	return nil
}
