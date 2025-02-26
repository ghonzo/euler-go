// Project Euler
package main

import "testing"

func Test_score(t *testing.T) {
	type args struct {
		name string
		pos  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{"COLIN", 938}, 49714},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := score(tt.args.name, tt.args.pos); got != tt.want {
				t.Errorf("score() = %v, want %v", got, tt.want)
			}
		})
	}
}
