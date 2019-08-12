package main

import (
	"testing"
)

// Tests are just regular Go functions with a few rules:
//      * The name of the test function must start with Test.
//      * The test function must take one argument of type *testing.T.
// A *testing.T is a type injected by the testing package itself,
// to provide ways to print, skip, and fail the test.

// https://dave.cheney.net/2019/05/07/prefer-table-driven-tests

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"main"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestGetHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"hello", "Hello, world"},
		// {"bad", "not this"},      // demonstrate a failing test; check that tests are running
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHello(); got != tt.want {
				t.Errorf("GetHello() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestHypotenuse(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"zero", args{0, 0}, 0},
		{"3-4-5", args{3, 4}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hypotenuse(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Hypotenuse() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestAverage(t *testing.T) {
	type args struct {
		values []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "one number",
			args: args{values: []float64{3.14}},
			want: 3.14},
		{
			name: "three numbers",
			args: args{values: []float64{3.0, 5.0, 7.0}},
			want: 5.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.args.values...); got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_newPrivateFuncThatIsNotTested(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		{"testing private function"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			newPrivateFuncThatIsNotTested()
// 		})
// 	}
// }
