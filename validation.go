package main

import (
	"fmt"
	"os"
)

func getFile() string {
	//reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose file name: ")
	fileName, _ := reader.ReadString('\n')
	fileName = fileName[0:len(fileName)-2] + ""

	return fileName
}

func exists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}
	return false
}

func checkRewrite(file string) bool {
	if exists(file) {
		//reader := bufio.NewReader(os.Stdin)
		fmt.Println(Teal("File " + file + " already exist, recreate? (print y)"))
		answer, _ := reader.ReadString('\n')
		answer = answer[0:len(answer)-2] + ""
		if answer != "y" {
			return false
		}
	}
	return true
}
