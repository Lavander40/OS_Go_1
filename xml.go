package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func showXmlMenu() {
	fmt.Println(Yellow("1) Show menu options"))
	fmt.Println(Yellow("2) Create xml"))
	fmt.Println(Yellow("3) Write in xml"))
	fmt.Println(Yellow("4) Read xml"))
	fmt.Println(Yellow("5) Delete xml"))
	fmt.Println(Yellow("6) Get back"))
}

func xmlMenu() {
	showXmlMenu()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Your choice: ")
		input, _ := reader.ReadString('\n')
		input = input[0:len(input)-2] + ""
		switch input {
		case "1":
			showXmlMenu()
		case "2":
			createXml()
		case "3":
			writeXml()
		case "4":
			readXml()
		case "5":
			deleteXml()
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

func createXml() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose file name: ")
	input, _ := reader.ReadString('\n')
	input = input[0:len(input)-2] + ""

	_, err := os.Stat(input)

	if err == nil {
		fmt.Println(Teal("File " + input + " already exist, rewrite? (print y)"))
		answer, _ := reader.ReadString('\n')
		answer = answer[0:len(answer)-2] + ""
		if answer != "y" {
			return
		}
	}

	file, err := os.Create("data/" + input + ".txt")

	if err != nil {
		fmt.Print(err)
		return
	}

	defer file.Close()
	fmt.Println(Green("File created"))
}

func writeXml() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose file name: ")
	input, _ := reader.ReadString('\n')
	input = input[0:len(input)-2] + ""

	file, err := os.Create("data/" + input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Print(Yellow("File was opened\nWrite text to input: "))
	input, _ = reader.ReadString('\n')

	_, err = file.WriteString(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Green("Writing complete"))
}

func readXml() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose file name: ")
	input, _ := reader.ReadString('\n')
	input = input[0:len(input)-2] + ""

	file, err := os.ReadFile("data/" + input)
	if err != nil {
		fmt.Print(Red(err))
		return
	}

	fmt.Println("File content:\n" + string(file))
}

func deleteXml() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose file name: ")
	input, _ := reader.ReadString('\n')
	input = input[0:len(input)-2] + ""

	if err := os.Remove("data/" + input); err != nil {
		fmt.Print(Red(err))
		return
	}

	fmt.Println(Green("File deleted"))
}
