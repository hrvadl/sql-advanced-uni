package main

import "testing"

func Test_removeDigits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ShouldRemove5And0",
			args: args{
				n: 59015509,
			},
			want: 919,
		},
		{
			name: "ShouldRemove5And0",
			args: args{
				n: 55001,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDigits(tt.args.n); got != tt.want {
				t.Errorf("removeDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
