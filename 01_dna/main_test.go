package main

import "testing"

func Test_countDNANucleotides(t *testing.T) {
	tests := []struct {
		name       string
		seq        string
		wantCountA int
		wantCountC int
		wantCountG int
		wantCountT int
	}{
		{
			"Test count DNA nucleotides",
			"AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC",
			20,
			12,
			17,
			21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCountA, gotCountC, gotCountG, gotCountT := countDNANucleotides(tt.seq)
			if gotCountA != tt.wantCountA {
				t.Errorf("countDNANucleotides() gotCountA = %v, want %v", gotCountA, tt.wantCountA)
			}
			if gotCountC != tt.wantCountC {
				t.Errorf("countDNANucleotides() gotCountC = %v, want %v", gotCountC, tt.wantCountC)
			}
			if gotCountG != tt.wantCountG {
				t.Errorf("countDNANucleotides() gotCountG = %v, want %v", gotCountG, tt.wantCountG)
			}
			if gotCountT != tt.wantCountT {
				t.Errorf("countDNANucleotides() gotCountT = %v, want %v", gotCountT, tt.wantCountT)
			}
		})
	}
}
