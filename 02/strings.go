package main

import "fmt"

func main() {
	firstName := "Chuck"
	familyName := "Norris"

	// fmt.Println(firstName)
	// fmt.Println(familyName)
	fmt.Println(firstName + " " + familyName)

	// Creating and initializing strings
	str1 := "Never"
	str2 := "the"
	str3 := "Machine"
	str4 := "Forever"

	// Concatenating strings using Sprintf() function
	result := fmt.Sprintf("%s %s %s %s!", str1, str2, str3, str4)

	fmt.Println(result)

}
