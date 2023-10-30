package main

import (
	"bufio"
	"fmt"
	"os"
)

func showFileMenu() {
	fmt.Println(Yellow("1) Show menu options"))
	fmt.Println(Yellow("2) Create file"))
	fmt.Println(Yellow("3) Write in file"))
	fmt.Println(Yellow("4) Read file"))
	fmt.Println(Yellow("5) Delete file"))
	fmt.Println(Yellow("6) Get back"))
}

func fileMenu() {
	showFileMenu()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Your choice: ")
		input, _ := reader.ReadString('\n')
		input = input[0:len(input)-2] + ""
		switch input {
		case "1":
			showFileMenu()
		case "2":
			createFile()
		case "3":
			writeFile()
		case "4":
			readFile()
		case "5":
			deleteFile()
		case "6":
			showMenu()
			return
		default:
			fmt.Println(Red("Invalid choice."))
			continue
		}
		//fmt.Println("You chose option " + Yellow((input))
	}
}

func createFile() {
	input := "data/" + getFile() + ".txt"

	if checkRewrite(input) == false {
		fmt.Println(Red("aborting operation"))
		return
	}

	file, err := os.Create(input)

	if err != nil {
		fmt.Print(Red("Creation error"))
		return
	}

	defer file.Close()
	fmt.Println(Green("File created"))
}

func writeFile() {
	input := "data/" + getFile() + ".txt"

	if checkRewrite(input) == false {
		fmt.Println(Yellow("aborting operation"))
		return
	}

	file, err := os.Create(input)
	if err != nil {
		fmt.Print(Red("Invalid input"))
		return
	}
	defer file.Close()

	//reader := bufio.NewReader(os.Stdin)
	fmt.Print(Yellow("File was opened\nWrite text to input: "))
	input, _ = reader.ReadString('\n')

	_, err = file.WriteString(input)
	if err != nil {
		fmt.Print(Red("Invalid input"))
		return
	}

	fmt.Println(Green("Writing complete"))
}

func readFile() {
	input := "data/" + getFile() + ".txt"

	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Print(Red("Invalid input"))
		return
	}

	fmt.Println(Yellow("File content:\n") + string(file))
}

func deleteFile() {
	input := "data/" + getFile() + ".txt"

	if err := os.Remove(input); err != nil {
		fmt.Print(Red("Invalid input"))
		return
	}

	fmt.Println(Green("File deleted"))
}
