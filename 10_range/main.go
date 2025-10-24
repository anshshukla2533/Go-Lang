package main

import "fmt"

func main() {
	/*
	nums := []int{6, 7, 8}
	var sum int;
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
		sum+=nums[i];git 
		
	}
	fmt.Println(sum)
   // sum1:=0;
	////for i,num:=range nums{
		fmt.Println(num,i)
		//sum1=sum1+num;
	}
	//fmt.Println(sum1)
	*/
	m:=map[string]string{"fname":"john","lname":"dosle"}
	for k,v:=range m{
		fmt.Println(k,v)
	}
	for i,c:=range "golang"{
		fmt.Println(i,c)
	}
}