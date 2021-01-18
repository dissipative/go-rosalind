package main

import "testing"

func Test_countGC(t *testing.T) {
	tests := []struct {
		name string
		seq  string
		want float64
	}{
		{
			name: "Rosalind_6404",
			seq:  "CCTGCGGAAGATCGGCACTAGAATAGCCAGAACCGTTTCTCTGAGGCTTCCGGCCTTCCCTCCCACTAATAATTCTGAGG",
			want: 53.750000,
		},
		{
			name: "Rosalind_5959",
			seq:  "CCATCGGTAGCGCATCCTTAGTCCAATTAAGTCCCTATCCAGGCGCTCCGCCGAAGGTCTATATCCATTTGTCAGCAGACACGC",
			want: 53.571429,
		},
		{
			name: "Rosalind_0808",
			seq:  "CCACCCTCGTGGTATGGCTAGGCATTCAGGAACCGGAGAACGCTTCAGACCAGCCCGGACTGGGAACCTGCGGGCAGTAGGTGGAAT",
			want: 60.919540,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGC(tt.seq); got != tt.want {
				t.Errorf("countGC() = %v, want %v", got, tt.want)
			}
		})
	}
}
