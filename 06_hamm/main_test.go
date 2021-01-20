package main

import "testing"

func Test_hamm(t *testing.T) {
	type args struct {
		seq1 string
		seq2 string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"Sample dataset",
			args{
				"GAGCCTACTAACGGGAT",
				"CATCGTAATGACGGCCT",
			},
			7,
			false,
		},
		{
			"Diff length error",
			args{
				"ATGC",
				"ATC",
			},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hamm(tt.args.seq1, tt.args.seq2)
			if (err != nil) != tt.wantErr {
				t.Errorf("hamm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hamm() = %v, want %v", got, tt.want)
			}
		})
	}
}
