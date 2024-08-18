package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

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

	// We can then call the WriteJSON method using a buffer - in memory - not ideal for prod.
	//though see how this satifies the interface
	var buf bytes.Buffer
	err := f.WriteJSON(&buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(&buf)

	// Or using a file. Same principle to showcase interface
	// to implement an interface you have to satisfy it by calling all
	// the methods defined in said interface.
	file, err := os.Create("/tmp/ferrets.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = f.WriteJSON(file)
	if err != nil {
		log.Fatal(err)
	}

}
