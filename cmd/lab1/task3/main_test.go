package main

import "testing"

func Test_gcd(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestGCD12And36",
			args: args{
				m: 12,
				n: 36,
			},
			want: 12,
		},
		{
			name: "TestGCD48And72",
			args: args{
				m: 48,
				n: 72,
			},
			want: 24,
		},
		{
			name: "TestGCD12And13",
			args: args{
				m: 12,
				n: 13,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}
