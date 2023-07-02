package main

import "fmt"

// Run Func would instantiate and startup the server

func Run() error {
	fmt.Println("Starting server...")
	return nil
}



func main() {
	fmt.Println("Hello world")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}