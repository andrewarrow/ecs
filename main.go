package main

import (
	"ecs/persist"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var school School
var argMap = map[string]string{}

func main() {
	rand.Seed(time.Now().UnixNano())
	persist.Init()
	fmt.Println("ecs")
	fmt.Println("")
	if len(os.Args) == 1 {
		fmt.Println("  ecs list          # list students at ecs")
		fmt.Println("  ecs add           # add a new student")
		fmt.Println("  ecs note          # add a note to a specific student")
		fmt.Println("")
		return
	}

	LoadSchool()
	argMap = argsToMap()
	command := os.Args[1]

	if command == "ls" || command == "list" {
		ListAll()
	} else if command == "add" {
		AddMenu()
	} else {
	}
}
