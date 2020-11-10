package main

import (
	"fmt"
	"time"
)

func ListHelp() {
	fmt.Println("")
}

func ListAll() {
	for _, s := range school.Students {
		dobString := fmt.Sprintf("%v", s.DateOfBirth)
		dobString = dobString[0:10]
		age := (time.Now().Unix() - s.DateOfBirth.Unix()) / 86400 / 365
		fmt.Printf("%03d. %20s, %20s  %20s (%d)\n", s.Id, s.LastName, s.FirstName, dobString, age)
	}
	fmt.Println("")
}
