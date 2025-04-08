package main

import (
	"reflect"
	"testing"
)

func Test_toDigits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "TestZero",
			args: args{
				n: 0,
			},
			want: []int{0},
		},
		{
			name: "TestSplit",
			args: args{
				n: 129078,
			},
			want: []int{1, 2, 9, 0, 7, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toDigits(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
