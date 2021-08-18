package leetcode

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t01", args{[]int{1, 2, 3, 4, 5, 6, 7}, 3}, []int{5, 6, 7, 1, 2, 3, 4}},
		{"t02", args{[]int{-1, -100, 3, 99}, 2}, []int{3, 99, -1, -100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			for i := range tt.args.nums {
				if tt.args.nums[i] != tt.want[i] {
					t.Errorf("unit test: %s, rotate() = %v, want %v\n", tt.name, tt.args.nums, tt.want)
				}
			}
		})
	}
}
