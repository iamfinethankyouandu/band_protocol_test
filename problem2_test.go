package main

import "testing"

func TestSupermanChickenRescue(t *testing.T) {
	type args struct {
		n        int
		k        int
		chickens []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				n:        5,
				k:        5,
				chickens: []int{2, 5, 10, 12, 15},
			},
			want: 2,
		},
		{
			args: args{
				n:        6,
				k:        10,
				chickens: []int{1, 11, 30, 34, 35, 37},
			},
			want: 4,
		},
		{
			args: args{
				n:        7,
				k:        11,
				chickens: []int{1, 11, 30, 34, 35, 37, 40},
			},
			want: 5,
		},
		{
			args: args{
				n:        3,
				k:        1,
				chickens: []int{9, 99, 999},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run("Problem 2: Superman's Chicken Rescue", func(t *testing.T) {
			if got := SupermanChickenRescue(tt.args.n, tt.args.k, tt.args.chickens); got != tt.want {
				t.Errorf("SupermanChickenRescue() = %v, want %v", got, tt.want)
			}
		})
	}
}
