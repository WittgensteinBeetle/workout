package main

import "fmt"

func main() {
	var fruit = []string{"dragonfruit", "pear", "watermelon"}

	fmt.Println(RemoveElement(fruit, "pear"))
	fmt.Println(Fibonacci(8))

	f := Business{}
	f.Init()

	//Iterate over the closed channel which is in a signaled state. Which makes all reads return immediately
	//this may be a useful idea for reading large chunks of data from files

	for ferret := range f.CollectionChannel() {
		fmt.Println("Channel values for Ferret Name,Age:", ferret.Name, ferret.Age)

	}
}
