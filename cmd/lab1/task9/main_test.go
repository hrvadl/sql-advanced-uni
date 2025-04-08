package main

import "testing"

func Test_process(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ShouldConvert2ToBinary",
			args: args{
				n: 2,
			},
			want: "10",
		},
		{
			name: "ShouldConvert3ToBinary",
			args: args{
				n: 3,
			},
			want: "11",
		},
		{
			name: "ShouldConvert8ToBinary",
			args: args{
				n: 8,
			},
			want: "1000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := process(tt.args.n); got != tt.want {
				t.Errorf("process() = %v, want %v", got, tt.want)
			}
		})
	}
}
