package twosum2

func twoSum(numbers []int, target int) []int {
	var sum int
	for i, j := 0, len(numbers)-1; i < j; {
		sum = numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}
		if sum > target {
			j = j - 1
		} else {
			i = i + 1
		}
	}
	return nil
}
