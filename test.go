package main

import "fmt"

const (
	Question_1 = (iota + 1)
	Question_2
	Question_3
)

const (
	YES = "YES"
	NO  = "NO"
)

func main() {
	for {
		fmt.Print("Type number of question:")
		var questionNumber int
		fmt.Scanln(&questionNumber)
		switch questionNumber {
		case Question_1:
			var returnedBook, expectedReturnedBook returnBook
			fmt.Scanln(&returnedBook.dd, &returnedBook.mm, &returnedBook.yy)
			fmt.Scanln(&expectedReturnedBook.dd, &expectedReturnedBook.mm, &expectedReturnedBook.yy)
			ans := test1(returnedBook, expectedReturnedBook)
			fmt.Println(ans)

		case Question_2:
			var sumStudent, sumCandie, start int64
			fmt.Scanln(&sumStudent, &sumCandie, &start)
			ans := test2(sumStudent, sumCandie, start)
			fmt.Println(ans)

		case Question_3:
			var n int
			fmt.Scanln(&n)
			arr := make([]int, n)
			for i := 0; i < n; i++ {
				fmt.Scan(&arr[i])
			}
			ans := test3(arr)
			fmt.Println(ans)
		default:
			return
		}
	}
}

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
			finalStudent = start
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
