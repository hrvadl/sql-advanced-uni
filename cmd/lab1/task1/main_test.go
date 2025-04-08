package main

import "testing"

func Test_process(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name           string
		args           args
		wantDigitCount int
		wantDigitSum   int
		wantFirstDigit int
	}{
		{
			name: "TestProcess12345",
			args: args{
				n: 12345,
			},
			wantDigitCount: 5,
			wantDigitSum:   15,
			wantFirstDigit: 1,
		},
		{
			name: "TestProcess333",
			args: args{
				n: 333,
			},
			wantDigitCount: 3,
			wantDigitSum:   9,
			wantFirstDigit: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDigitCount, gotDigitSum, gotFirstDigit := process(tt.args.n)
			if gotDigitCount != tt.wantDigitCount {
				t.Errorf("process() gotDigitCount = %v, want %v", gotDigitCount, tt.wantDigitCount)
			}
			if gotDigitSum != tt.wantDigitSum {
				t.Errorf("process() gotDigitSum = %v, want %v", gotDigitSum, tt.wantDigitSum)
			}
			if gotFirstDigit != tt.wantFirstDigit {
				t.Errorf("process() gotFirstDigit = %v, want %v", gotFirstDigit, tt.wantFirstDigit)
			}
		})
	}
}
