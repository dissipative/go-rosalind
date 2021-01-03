package main

import "testing"

func Test_rna(t *testing.T) {

	tests := []struct {
		name string
		seq  string
		want string
	}{
		{
			"Test DNA to RNA conversion",
			"GATGGAACTTGACTACGTAAATT",
			"GAUGGAACUUGACUACGUAAAUU",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rna(tt.seq); got != tt.want {
				t.Errorf("rna() = %v, want %v", got, tt.want)
			}
		})
	}
}
