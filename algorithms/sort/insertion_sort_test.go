package sortalgo

import (
	"reflect"
	"testing"
)

// func TestInsertionSort(t *testing.T) {
// 	tests := []struct {
// 		input  []int
// 		expect []int
// 	}{
// 		{[]int{3, 2, 1}, []int{1, 2, 3}},
// 		{[]int{}, []int{}},
// 	}

// 	for _, test := range tests {
// 		got := insertionSort(test.input)
// 		for i := range test.expect {
// 			if test.expect[i] != got[i] {
// 				t.Errorf("expect %v, but got %v\n", test.expect, got)
// 				break
// 			}
// 		}
// 	}
// }

func Test_insertionSort2(t *testing.T) {
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
			if got := insertionSort2(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSort2() = %v, want %v", got, tt.want)
			}
		})
	}
}
