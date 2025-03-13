package main
import "fmt"
func main(){
	grade :=map[string]int{
		"Najim":95,
		"Suhail":91,
		"Muzzamil":90,
		"Farseen":85,

	}
	fmt.Println("Initial Grades :", grade)

	grade["Suhail"] = 88

	grade["Midulaj"] = 75

	fmt.Println("Update Grades:", grade)
	
	student := "Njaim"
	fmt.Println(student+ " 'S Grades:",grade[student])

	student = " Abdul sathar "
	if _,exists :=grade[student];exists{
		grade[student] = 90
	}else{
		fmt.Println(student,"Not found in the records.")
	}
}