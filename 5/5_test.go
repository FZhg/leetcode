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
			if got := brutalForceLongestPalindromeSubstring(tt.args.s); got != tt.want {
				t.Errorf("brutalForceLongestPalindromeSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateStubString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"TestCase",
			args{
				s: "ababac",
			},
			"@#a#b#a#b#a#c#$",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateStubString(tt.args.s); got != tt.want {
				t.Errorf("generateStubString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_manacharLongestPalindromeSubstring(t *testing.T) {
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
			if got := manacharLongestPalindromeSubstring(tt.args.s); got != tt.want {
				t.Errorf("manacharLongestPalindromeSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
