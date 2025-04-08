package main

import (
	"reflect"
	"testing"
)

func Test_findDivisors(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "TestDivisorsOf36",
			args: args{
				input: 36,
			},
			want: []int{1, 2, 3, 4, 6, 9, 12, 18, 36},
		},
		{
			name: "TestDivisorsOf36",
			args: args{
				input: 13,
			},
			want: []int{1, 13},
		},
		{
			name: "TestDivisorsOf39",
			args: args{
				input: 39,
			},
			want: []int{1, 3, 13, 39},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDivisors(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
