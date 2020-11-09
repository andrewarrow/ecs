package main

import (
	"fmt"
)

func ListHelp() {
	fmt.Println("")
}

func ListAll() {
	for i, s := range school.Students {
		dobString := fmt.Sprintf("%v", s.DateOfBirth)
		dobString = dobString[0:10]
		fmt.Printf("%03d. %20s, %20s  %20s\n", i+1, s.LastName, s.FirstName, dobString)
	}
	fmt.Println("")
}
