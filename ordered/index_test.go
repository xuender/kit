// nolint: paralleltest
package ordered_test

import (
	"testing"

	"github.com/xuender/kit/ordered"
)

type args struct {
	slice []int
	elem  int
}

func TestIndexAes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args args
		want int
	}{
		{"nil", args{[]int{}, 1}, 0},
		{"1", args{[]int{1}, 2}, 1},
		{"1", args{[]int{2}, 1}, 0},
		{"2", args{[]int{2, 3}, 1}, 0},
		{"2", args{[]int{2, 3}, 4}, 2},
		{"3", args{[]int{2, 4}, 3}, 1},
		{"3", args{[]int{1, 1, 2, 4}, 3}, 3},
		{"3", args{[]int{1, 2, 4}, 3}, 2},
	}

	for _, data := range tests {
		t.Run(data.name, func(t *testing.T) {
			if got := ordered.IndexAes(data.args.slice, data.args.elem); got != data.want {
				t.Errorf("IndexAes() = %v, want %v", got, data.want)
			}
		})
	}
}

func TestIndexSet(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args args
		want int
	}{
		{"nil", args{[]int{}, 1}, 0},
		{"1", args{[]int{1}, 2}, 1},
		{"1", args{[]int{2}, 1}, 0},
		{"-1", args{[]int{1}, 1}, -1},
		{"1", args{[]int{2}, 1}, 0},
		{"2", args{[]int{2, 3}, 1}, 0},
		{"2", args{[]int{2, 3}, 4}, 2},
		{"3", args{[]int{2, 4}, 3}, 1},
		{"3", args{[]int{1, 2, 4}, 3}, 2},
		{"3", args{[]int{1, 2, 4}, 3}, 2},
		{"-3", args{[]int{1, 2, 4}, 2}, -1},
	}

	for _, data := range tests {
		t.Run(data.name, func(t *testing.T) {
			if got := ordered.IndexSet(data.args.slice, data.args.elem); got != data.want {
				t.Errorf("IndexAes() = %v, want %v", got, data.want)
			}
		})
	}
}

func TestIndexDesc(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args args
		want int
	}{
		{"nil", args{[]int{}, 1}, 0},
		{"1", args{[]int{1}, 2}, 0},
		{"1", args{[]int{2}, 1}, 1},
		{"2", args{[]int{3, 2}, 1}, 2},
		{"2", args{[]int{3, 2}, 4}, 0},
		{"3", args{[]int{4, 2}, 3}, 1},
		{"3", args{[]int{4, 2, 1, 1}, 3}, 1},
		{"3", args{[]int{5, 5, 4, 2}, 3}, 3},
		{"3", args{[]int{4, 2, 1}, 3}, 1},
	}

	for _, data := range tests {
		t.Run(data.name, func(t *testing.T) {
			if got := ordered.IndexDesc(data.args.slice, data.args.elem); got != data.want {
				t.Errorf("IndexAes() = %v, want %v", got, data.want)
			}
		})
	}
}
