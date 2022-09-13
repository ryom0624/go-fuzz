package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		// TODO: Add test cases.
		{"Hello World", "dlroW olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := Reverse(tt.in)
			if err != nil {
				return
			}
			if got != tt.want {
				t.Errorf("Reverse() = %q, want %q", got, tt.want)
			}
		})
	}
}

/*
	1. Reversing a string twice preserves the original value
	2. The reversed string preserves its state as valid UTF-8.
*/
func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, World", " ", "!12345"}

	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
			//t.Skip()
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
			//t.Skip()
		}
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}

		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
