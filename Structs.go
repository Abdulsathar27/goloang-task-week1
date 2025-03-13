package main
import "fmt"
type Person struct{
	Name string
	Age int 
}

func (p Person)PrintDetails(){
	fmt.Printf("Name:%s,Age:%d\n",p.Name,p.Age)
}

func main(){
	Person1 := Person {Name:"Alice",Age:25}

	Person1.PrintDetails()
}