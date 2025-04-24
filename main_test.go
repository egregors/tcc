package main

import "testing"

func Test_solve(t *testing.T) {
	tests := []struct {
		name      string
		inputPath string
		want      int
	}{
		{
			name:      "test1",
			inputPath: "./test.txt",
			want:      15,
		},
		{
			name:      "test2",
			inputPath: "./test2.txt",
			want:      15,
		},
		{
			name:      "test3",
			inputPath: "./test3.txt",
			want:      12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.inputPath); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
