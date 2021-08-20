package leetcode

import "testing"

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t01", args{[]int{0, 1, 0, 3, 12}}, []int{1, 3, 12, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			for i := 0; i < len(tt.args.nums); i++ {
				if tt.args.nums[i] != tt.want[i] {
					t.Errorf("ut: %s, moveZeroes() = %v, want = %v\n", tt.name, tt.args.nums, tt.want)
				}
			}
		})
	}
}
