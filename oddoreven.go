package main
import "fmt"
func main(){
	var Number int
	//Ask number from user
	fmt.Println("Enter a value : ")
	_, err := fmt.Scan(&Number)

	if err!=nil{
		fmt.Println("Please Enter an integer.")
		return 
	}
	//check if the number is even or odd
	if Number%2 == 0{
		fmt.Println(Number,"is even.")
	}else{
		fmt.Println(Number,"is odd.")
	}
}


