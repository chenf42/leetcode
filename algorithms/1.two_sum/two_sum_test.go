package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{1, 1}, 2, []int{0, 1}},
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for _, testCase := range testCases {
		got := twoSum(testCase.nums, testCase.target)
		if !reflect.DeepEqual(got, testCase.expected) {
			t.Errorf("Expected testCase(%v, %v) == %v, but got %v", testCase.nums, testCase.target, testCase.expected, got)
		}
	}
}
