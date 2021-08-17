package leetcode

import (
	"fmt"
	"testing"
)

func genArray(n int, bad int) []bool {
	arr := make([]bool, n)
	for i := 0; i < bad-1; i++ {
		arr[i] = true
	}

	return arr
}

func Test_firstBadVersion(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01", args{5}, 4},
		{"test02", args{1}, 1},
		{"test03", args{5}, 3},
	}
	for _, tt := range tests {
		arr = genArray(tt.args.n, tt.want)
		fmt.Printf("arr: %v\n", arr)

		t.Run(tt.name, func(t *testing.T) {
			if got := firstBadVersion(tt.args.n); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
