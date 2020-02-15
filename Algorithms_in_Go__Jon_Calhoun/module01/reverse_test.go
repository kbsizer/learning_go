package module01

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	var nilString string
	tests := []struct {
		name string
		word string
		want string
	}{
		{"nil", nilString, nilString},
		{"empty string", "", ""},
		{"3 letter word", "cat", "tac"},
		{"8 letter word", "alphabet", "tebahpla"},
		{"non-ASCII word", "日本語", "語本日"},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.name), func(t *testing.T) {
			got := Reverse(tc.word)
			if got != tc.want {
				t.Fatalf("Reverse() = %v; want %v", got, tc.want)
			}
		})
	}
}
