package main

import "fmt"
//similar to dictionary
func main() {
	m := make(map[string]string)
	m["name"] = "golang"
	m["area"]="backend"
	fmt.Println(m["name"],m["area"])
}