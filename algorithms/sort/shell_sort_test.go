package sortalgo

import (
	"reflect"
	"testing"
)

func Test_shellSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"unit test 01", args{[]int{8, 6, 3, 4, 2, 1}}, []int{1, 2, 3, 4, 6, 8}},
		{"unit test 01", args{[]int{3, 2, 1}}, []int{1, 2, 3}},
		{"unit test 02", args{[]int{}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shellSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shellSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
