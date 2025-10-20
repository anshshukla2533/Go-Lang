package main

import "fmt"
const age=40
//name:="ansh" not allwoed
func main() {
	const name string = "ansh"
	
	fmt.Println(age)
	fmt.Println(name)
	const (
		port=5000
		host="localhost"

	)
	fmt.Println(port,host);
}