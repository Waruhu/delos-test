package main

import (
	"fmt"
	"math"
)

const (
	Question_1 = (iota + 1)
	Question_2
	Question_3
)

const (
	YES = "YES"
	NO  = "NO"
)

type returnBook struct {
	dd, mm, yy int64
}

func test1(returnedBook, expectedReturnedBook returnBook) int64 {
	ans := int64(0)
	diffDay := int64(15)
	diffMonth := int64(500)
	diffYear := int64(12000)

	if returnedBook.yy != expectedReturnedBook.yy {
		ans = diffYear * (expectedReturnedBook.yy - returnedBook.yy)
	} else if returnedBook.mm != expectedReturnedBook.mm {
		ans = diffMonth * (expectedReturnedBook.mm - returnedBook.mm)
	} else {
		ans = diffDay * (expectedReturnedBook.dd - returnedBook.dd)
	}

	return ans
}

func test2(sumStudent, sumCandie, start int64) int64 {
	indexCandie := int64(1)
	finalStudent := int64(1)
	for indexCandie <= sumCandie {
		if indexCandie == sumCandie {
			finalStudent = start
			break
		}

		if start == sumStudent {
			start = 0
		}
		start++
		indexCandie++
	}

	return finalStudent
}

func test3(arr []int) string {
	if len(arr) == 1 {
		return YES
	}

	var totalLeft, totalRight int
	for i := 1; i < len(arr); i++ {
		totalRight += arr[i]
	}

	for i, j := 0, 1; j < len(arr); {
		totalRight -= arr[j]
		totalLeft += arr[i]

		if totalLeft == totalRight {
			return YES
		}

		i++
		j++
	}

	return NO
}

type matrix struct {
	arr [][]int
}

func main() {
	matrix := matrix{
		[][]int{
			{1, 4, 5, 1},
			{3, 6, 9, 1},
			{7, 1, 3, 1},
			{2, 2, 2, 2},
		},
	}
	fmt.Println(findAbs(matrix))
}

func findAbs(arr matrix) float64 {
	n := len(arr.arr)
	var diag1, step1, step2, diag2 int
	step1, step2 = 0, n-1

	for j := 0; j < n; j++ {
		diag1 += arr.arr[step1][j]
		step1++
		diag2 += arr.arr[j][step2]
		step2--
	}
	abs := diag1 - diag2
	return math.Abs(float64(abs))

}
