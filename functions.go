package main

import "fmt"

func RemoveElement(s []string, e string) []string {

	//condition for element in the begining and middle of the slice
	for i := 0; i < len(s); i++ {
		if s[i] == e {
			println("I found this in the slice")
			s = append(s[:i], s[i+1:]...)
			println("I removed the element: ", e)
		}
	}

	return s
}

func Fibonacci(n int) int {
	//slice length and capacity
	f := make([]int, n+1, n+2)

	//Input sanitization to make sure the passed value is a positive
	if n < 0 {
		fmt.Println("Negative numbers are not valid")
		return 0
	}

	//Should the number be less than 2, 0 is one that showcases this logic
	//then set the length of the slice to 2
	if n < 2 {
		//Setting length of f to 2 and capacity 2
		f = f[0:2]
	}

	//first two elements of the Fib sequence
	f[0] = 0
	f[1] = 1

	//Fib sequence
	//example 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 …

	//Idea: make this a while loop with break condition on certain numbers
	//else limit range
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
