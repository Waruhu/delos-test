package main

import (
	"testing"
)

func Test_test1(t *testing.T) {
	type args struct {
		returnedBook         returnBook
		expectedReturnedBook returnBook
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test case 1: diffrent day",
			args: args{
				returnedBook: returnBook{
					dd: 15,
					mm: 8,
					yy: 2022,
				},
				expectedReturnedBook: returnBook{
					dd: 22,
					mm: 8,
					yy: 2022,
				},
			},
			want: 105,
		},
		{
			name: "test case 2: diffrent month",
			args: args{
				returnedBook: returnBook{
					dd: 7,
					mm: 6,
					yy: 2022,
				},
				expectedReturnedBook: returnBook{
					dd: 19,
					mm: 8,
					yy: 2022,
				},
			},
			want: 1000,
		},
		{
			name: "test case 3: diffrent year",
			args: args{
				returnedBook: returnBook{
					dd: 7,
					mm: 6,
					yy: 2022,
				},
				expectedReturnedBook: returnBook{
					dd: 19,
					mm: 8,
					yy: 2023,
				},
			},
			want: 12000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := test1(tt.args.returnedBook, tt.args.expectedReturnedBook); got != tt.want {
				t.Errorf("test1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_test2(t *testing.T) {
	type args struct {
		sumStudent int64
		sumCandie  int64
		start      int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test case 1",
			args: args{
				sumStudent: 3,
				sumCandie:  5,
				start:      2,
			},
			want: 3,
		},
		{
			name: "test case 2",
			args: args{
				sumStudent: 5,
				sumCandie:  3,
				start:      1,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := test2(tt.args.sumStudent, tt.args.sumCandie, tt.args.start); got != tt.want {
				t.Errorf("test2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_test3(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{
				arr: []int{1, 3, 5, 4},
			},
			want: YES,
		},
		{
			name: "test case 2",
			args: args{
				arr: []int{1, 5, 7, 2, 4},
			},
			want: YES,
		},
		{
			name: "test case 3",
			args: args{
				arr: []int{1, 3, 4, 9},
			},
			want: NO,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := test3(tt.args.arr); got != tt.want {
				t.Errorf("test3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findAbs(t *testing.T) {
	type args struct {
		arr matrix
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test case 1",
			args: args{
				arr: matrix{
					arr: [][]int{
						{1, 4, 5, 1},
						{3, 6, 9, 1},
						{7, 1, 3, 1},
						{2, 2, 2, 2},
					},
				},
			},
			want: 1,
		},
		{
			name: "test case 2",
			args: args{
				arr: matrix{
					arr: [][]int{
						{1, 4, 5},
						{3, 6, 9},
						{7, 1, 3},
					},
				},
			},
			want: 8,
		},
		{
			name: "test case 3",
			args: args{
				arr: matrix{
					arr: [][]int{
						{4, 4, 4, 4},
						{4, 4, 4, 4},
						{4, 4, 4, 4},
						{4, 4, 4, 4},
					},
				},
			},
			want: 0,
		},
		{
			name: "test case 4",
			args: args{
				arr: matrix{
					arr: [][]int{
						{1, 2, 3, 4},
						{5, 6, 7, 8},
						{9, 10, 11, 12},
						{13, 14, 15, 16},
					},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAbs(tt.args.arr); got != tt.want {
				t.Errorf("findAbs() = %v, want %v", got, tt.want)
			}
		})
	}
}
