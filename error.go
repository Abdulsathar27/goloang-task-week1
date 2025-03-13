package main
import (
	"fmt"
"errors"
)
// Function to perform division with error handling
func divide(a,b float64)(float64,error){
	if b == 0{
		return 0,errors.New("division by zero is not allowed")
	}
	return a/b,nil
}

func main (){
	a:=12.00
	b:= 2.00
	result,err:=divide(a,b)
	if err!=nil{
		fmt.Println("Error:",err)
	}else{
		fmt.Println("Result:",result)
	}
}