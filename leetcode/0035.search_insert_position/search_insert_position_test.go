package leetcode

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t01", args{[]int{1, 3, 5, 6}, 5}, 2},
		{"t02", args{[]int{1, 3, 5, 6}, 2}, 1},
		{"t03", args{[]int{1, 3, 5, 6}, 7}, 4},
		{"t04", args{[]int{1, 3, 5, 6}, 0}, 0},
		{"t05", args{[]int{1}, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("unit test: %s, searchInsert() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
