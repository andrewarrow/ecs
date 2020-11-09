package main

import (
	"fmt"
	"os"
	"time"
)

func AddHelp() {
	fmt.Println("  ecs add --fname=Frank --lname=Smith --dob=2013-01-14")
	fmt.Println("")
}

func AddMenu() {
	if len(os.Args) == 2 {
		AddHelp()
		return
	}
	if os.Args[1] == "add" {
		AddStudent()
		return
	}
}

func AddStudent() {
	fname := argMap["fname"]
	lname := argMap["lname"]
	dob := argMap["dob"]
	if fname == "" || lname == "" || dob == "" {
		AddHelp()
		return
	}

	layout := "2006-01-02"
	dobTime, _ := time.Parse(layout, dob)
	s := Student{}
	s.FirstName = fname
	s.LastName = lname
	s.DateOfBirth = dobTime

	fmt.Println(s)
	school.Students = append(school.Students, s)
	SaveSchool()
}
