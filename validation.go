package main

import (
	"bufio"
	"fmt"
	"os"
)

func getFile() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose file name: ")
	fileName, _ := reader.ReadString('\n')
	fileName = fileName[0:len(fileName)-2] + ""

	return fileName
}

func exists() bool {
	if _, err := os.Stat("/path/to/whatever"); err == nil {
		return true
	}
	return false
}

func rewrite() {

}
