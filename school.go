package main

import (
	"ecs/persist"
	"encoding/json"
	"time"
)

type School struct {
	Students []Student `json:"students"`
}

type Student struct {
	FirstName   string    `json:"fname"`
	LastName    string    `json:"lname"`
	DateOfBirth time.Time `json:"dob"`
}

func LoadSchool() {
	jsonString := persist.ReadFromFile("school.json")
	//fmt.Println(jsonString)
	if jsonString == "" {
		return
	}
	json.Unmarshal([]byte(jsonString), &school)
}

func SaveSchool() {
	asBytes, _ := json.Marshal(school)
	persist.SaveToFile("school.json", string(asBytes))
}
