package main 
import "fmt"

//funcation to sovle sum,difference,product,and quotient
func calculate(a,b float64) (float64,float64,float64,float64){
	sum := a + b
	difference := a - b
	product := a * b
	quotient := a / b // b is never zero

	return sum,difference,product,quotient
}

func main (){
	a,b := 11.0,6.0
	sum,difference,product,quotient := calculate(a,b)

	fmt.Println("sum:",sum)
	fmt.Println("difference:",difference)
	fmt.Println("product:",product)
	fmt.Println("quotinet:",quotient)
}