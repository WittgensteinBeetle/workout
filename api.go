package main

type Ferret struct {
	Name string
	Age  int
}

// Business struc contains a slice of pointers named Data to the Ferret struc
type Business struct {
	Data []*Ferret
}
