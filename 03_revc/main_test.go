package main

import "testing"

func Test_revc(t *testing.T) {
	tests := []struct {
		name        string
		seq         string
		wantCompSeq string
	}{
		{
			"Test reverse complementarity function",
			"AAAACCCGGT",
			"ACCGGGTTTT",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCompSeq := revc(tt.seq); gotCompSeq != tt.wantCompSeq {
				t.Errorf("revc() = %v, want %v", gotCompSeq, tt.wantCompSeq)
			}
		})
	}
}
