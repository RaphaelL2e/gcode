package utils

import (
	"reflect"
	"testing"

	. "gcode/model"
)

func TestArrayToListNode(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				arr: []int{3, 2, 0, 4},
			},
			want: ArrayToListNode([]int{3, 2, 0, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayToListNode(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayToListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNodeToArray(t *testing.T) {
	type args struct {
		node *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				node: ArrayToListNode([]int{3, 2, 0, 4}),
			},
			want: []int{3, 2, 0, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListNodeToArray(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListNodeToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
