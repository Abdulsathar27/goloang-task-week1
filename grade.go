package main

import "fmt"

func main() {
	
	grades := map[string]int{
		"Najim":     95,
		"Suhail":    91,
		"Muzzamil":  90,
		"Farseen":   85,
	}

	
	fmt.Println("Initial Grades:", grades)

	
	grades["Suhail"] = 88

	grades["Midulaj"] = 75


	fmt.Println("Updated Grades:", grades)

	
	student := "Najim"
	fmt.Printf("%s's Grade: %d\n", student, grades[student])

	student = "Abdul Sathar"
	if _, exists := grades[student]; exists {
		grades[student] = 90 
	} else {
		fmt.Println(student, "not found in the records.")
	}
}
