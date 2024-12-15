package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int = &age

	fmt.Println("Age:", *agePointer)
	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)

	// editAgeToAdultYears(agePointer)
	// fmt.Println(age)
}

func getAdultYears(age *int) int {
	return *age - 18
}

// func editAgeToAdultYears(age *int) {
// 	*age = *age - 18
// }
