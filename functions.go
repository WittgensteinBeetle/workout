package main

import (
	"encoding/json"
	"fmt"
	"io"
)

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

// The reciver b of the Business struc pointer has a method named Init
func (b *Business) Init() {

	//Initialize struc pointers of each ferret, we do not want to make a copy
	wallace := &Ferret{"Wallace", 1}
	gromit := &Ferret{"Gromit", 2}
	eve := &Ferret{"Eve", 3}

	//return a slice of pointers
	b.Data = []*Ferret{wallace, gromit, eve}
}

// the reciever b of the Business struc has a method named CollectionChannel()
// which will return pointers of the channel of type Ferret struc
func (b *Business) CollectionChannel() chan *Ferret {

	//create a chanel of type Ferret struc with the length of the passed receiver Data
	//this ia a buffered channel, so it does not need a channel receiver
	//and the buffer is set to the length of the returned slice
	dataChannel := make(chan *Ferret, len(b.Data))

	for _, ferret := range b.Data {
		dataChannel <- ferret
	}

	close(dataChannel)
	return dataChannel
}

// Implement a WriteJSON method that takes an io.Writer as the parameter.
// It marshals the Ferret struct to JSON, and if the marshal worked
// successfully, then calls the relevant io.Writer's Write() method.

// WriteJSON Note the method is declared with a pointer receiver, which will automatically dereference,
// the pointer when using dot operator on it to access a field else call a method
func (c *Business) WriteJSON(w io.Writer) error {
	js, err := json.Marshal(c)
	if err != nil {
		return err
	}

	_, err = w.Write(js)
	return err
}
