package main

import "fmt"

var (
	Red    = color("\033[1;31m%s\033[0m")
	Green  = color("\033[1;32m%s\033[0m")
	Yellow = color("\033[1;33m%s\033[0m")
	Teal   = color("\033[1;36m%s\033[0m")
)

func color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}
