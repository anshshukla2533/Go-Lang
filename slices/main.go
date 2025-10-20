package main

import (
	"fmt"
	//"slices"
)

//automactilly dynamacyy memory keep increaing
//most used constrcut in go
//+ useful emthod on array
func main(){
 //unitillized slice is nil
  //var nums []int 

  //fmt.Println(len(nums))
  //var nums=make([]int,0,5)
  //capacity->maximum number of elements can fit
  //nums=append(nums, 4)
  //nums=append(nums, 5)
 // nums=append(nums, 6)
 // nums=append(nums, 7)
 //nums[0]=1
 // fmt.Println(nums)
 // fmt.Println(cap(nums))
  //fmt.Println(nums==nil)
  //copy fucntion
  //var nums1=make([]int,0,3)
  // nums1=append(nums1,5);
 // var nums2=make([]int,len(nums1))
 
  //copying
  //copy(nums2,nums1)
  //fmt.Println(nums1,nums2)
  //slice operator
 // var nums=[]int{1,2,3}
  //fmt.Println(nums[0:1])
  //slice pakcage
  //var nums1=[]int{1,2}
  //var nums2=[]int{1,2,3}
  //fmt.Println(slices.Equal(nums1,nums2))

  var nums=[][]int{{1,2,3,4},{1,2,3,4}}
  fmt.Println((nums))


}