package main

import (
	"fmt"
	"maps"
)

//similar to dictionary
func main() {
	/*
	m := make(map[string]string)
	m["name"] = "golang"
	m["area"]="backend"
	fmt.Println(m["name"],m["area"])
	*/
	/*m:=make(map[string]int)
	m["age"]=60;
	m["price"]=45;
	clear(m) 
	fmt.Println(m["age"])
	delete(m,"price") //to delete
	//to empty whole map
	*/
	m:=map[string]int{"price":40,"phones":3}
	m2:=map[string]int{"price":40,"phones":3}
	v,ok:= m["phones"]
	fmt.Println(v)
	if !ok{
		fmt.Println("all ok")
	} else{
		fmt.Println("not ok")
	}
	fmt.Println(m)
	fmt.Println(maps.Equal(m,m2))

}