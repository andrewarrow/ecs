package main

import (
	"os"
	"strings"
)

func argsToMap() map[string]string {
	m := map[string]string{}
	if len(os.Args) == 1 {
		return m
	}

	for _, a := range os.Args[1:] {
		if strings.HasPrefix(a, "--") {
			tokens := strings.Split(a, "=")
			if len(tokens) == 2 {
				key := strings.Split(tokens[0], "--")
				m[key[1]] = tokens[1]
			}
		}
	}
	return m
}
