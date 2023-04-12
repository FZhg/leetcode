package _5

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"testcase 1",
			args{
				s: "babad",
			},
			"bab",
		},
		{
			"testcase 2",
			args{
				s: "cbbd",
			},
			"bb",
		},
		{
			"testcase 3",
			args{
				s: "cbbb",
			},
			"bbb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
