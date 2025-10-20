package main

import "fmt"

//switch
//no need of writing beaark in go ,it alwyas in underhood
//multiple condition switch

func main(){
	/*
  i:=3;
  switch i{
  case 1:
	fmt.Println("one")
  case 2:
	fmt.Println("two")
  case 3:
	fmt.Println("three")
  default:
	fmt.Println("other")
   
  }
  switch time.Now().Weekday(){
  case time.Saturday,time.Monday:
	fmt.Println("its a weekend")
  default:
	fmt.Println("its workday")
  }
	*/
	//type switch

  whoAmI:=func (i interface{})  {
	switch t :=i.(type){
	case int:
		fmt.Println("integer")
	case string:
		fmt.Println("hey its string")
	case bool:
		fmt.Println("its",t)
	}
  }

  whoAmI(true)

}