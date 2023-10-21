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
	input := getFile()

	_, err := os.Stat("data/" + input + ".txt")

	if err == nil {
		reader := bufio.NewReader(os.Stdin)

		if _, err := os.Stat("/path/to/whatever"); err == nil {
			fmt.Println(Teal("File " + input + " already exist, recreate? (print y)"))
			answer, _ := reader.ReadString('\n')
			answer = answer[0:len(answer)-2] + ""
			if answer != "y" {
				return
			}
		}
	}

	file, err := os.Create("data/" + input + ".txt")

	if err != nil {
		fmt.Print(Red("Creation error"))
		return
	}

	defer file.Close()
	fmt.Println(Green("File created"))
}

func writeFile() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose file name: ")
	input, _ := reader.ReadString('\n')
	input = input[0:len(input)-2] + ""

	file, err := os.Create("data/" + input + ".txt")
	if err != nil {
		fmt.Print(Red("Invalid input"))
		return
	}
	defer file.Close()

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
	input := getFile()

	file, err := os.ReadFile("data/" + input + ".txt")
	if err != nil {
		fmt.Print(Red("Invalid input"))
		return
	}

	fmt.Println("File content:\n" + string(file))
}

func deleteFile() {
	input := getFile()

	if err := os.Remove("data/" + input + ".txt"); err != nil {
		fmt.Print(Red("Invalid input"))
		return
	}

	fmt.Println(Green("File deleted"))
}
