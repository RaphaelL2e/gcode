package set_matrix_zeroes

import (
	"reflect"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
			ok := reflect.DeepEqual(tt.args.matrix, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}})
			if !ok {
				t.Errorf("failed matrix: %v", tt.args.matrix)
			}
		})
	}
}
