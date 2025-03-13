package main
import (
	"fmt"
)
  func findMaxMin (arr []int)(int,int){
	max:=arr[0]
	min:=arr[0]
	for _,value :=range arr{
		if value>max{
			max=value
		}
		if value < min{
			min=value
		}

	}
	return max ,min
  }
  func main (){
	array := []int{10,6,9,7,5,3,4,8,1,2}
	max,min:=findMaxMin(array)
	fmt.Println("Maximum",max)
	fmt.Println("Minimum",min)
  }
