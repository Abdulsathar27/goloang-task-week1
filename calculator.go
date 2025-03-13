package main
import "fmt"
func main(){

     // simple calaculation function  
	fmt.Println("CALCULATOR")
	fmt.Println("-------------MENU-------------")
	fmt.Println("1.Add")
	fmt.Println("2.Subtract")
	fmt.Println("3.Multiply")
	fmt.Println("4.Divide")
	fmt.Println("------------------------------")
	fmt.Println("Choose The Operation")

	//user choice are store
	var choice int

	//as per choice shown 
	fmt.Scan(&choice)
	var a,b int

	// user choose the operation and give 2 two number 
	
	fmt.Println("Enter The Number a: ")
	fmt.Scan(&a)
	fmt.Println("Enter The Number b: ")
	fmt.Scan(&b)

	//user chooseing operation part 1 @ addition
	if(choice==1){
		ans:= a+b
		fmt.Println("=",ans)
	}else if(choice==2){
		ans := a-b
		fmt.Println("=",ans)
	}else if(choice==3){
		ans := a*b
		fmt.Println("=",ans)
	}else{
		ans := a/b
		fmt.Println("=",ans)
	}

	fmt.Println("-------------------------------")
	fmt.Println("Thank you !,Have a nice day ")
	fmt.Println("-------------------------------")

	
}