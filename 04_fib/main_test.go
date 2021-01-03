package main

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test reccurent",
			args{5, 3},
			19,
		},
		{
			"Test reccurent with another k",
			args{5, 2},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reccurent(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
