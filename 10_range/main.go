package main

import "fmt"

func main() {
	nums := []int{6, 7, 8}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}