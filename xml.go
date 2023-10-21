package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
)

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Age       int      `xml:"age"`
}

func NewPerson(firstName, lastName string) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       23,
	}
}

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

	_, err := os.Stat("data/" + input + ".xml")

	if err == nil {
		fmt.Println(Teal("File " + input + " already exist, rewrite? (print y)"))
		answer, _ := reader.ReadString('\n')
		answer = answer[0:len(answer)-2] + ""
		if answer != "y" {
			return
		}
	}

	file, err := os.Create("data/" + input + ".xml")
	if err != nil {
		fmt.Print(Red("File creation error"))
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

	file, err := os.Create("data/" + input + ".xml")
	if err != nil {
		fmt.Print(Red("Creation error"))
		return
	}
	defer file.Close()

	fmt.Println("File was opened\nCreating User instance...")
	fmt.Println(Yellow("Input first and last names:"))
	firstName, _ := reader.ReadString('\n')
	firstName = firstName[0:len(firstName)-2] + ""
	lastName, _ := reader.ReadString('\n')
	lastName = lastName[0:len(lastName)-2] + ""

	person := NewPerson(firstName, lastName)
	output, err := xml.MarshalIndent(person, "", " ")
	_, err = file.WriteString(xml.Header)
	_, err = file.Write(output)
	if err != nil {
		fmt.Print(Red("Writing error"))
		return
	}

	fmt.Println(Green("Writing complete"))
}

func readXml() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose file name: ")
	input, _ := reader.ReadString('\n')
	input = input[0:len(input)-2] + ""

	file, err := os.ReadFile("data/" + input + ".xml")
	if err != nil {
		fmt.Printf(Red("File %s does not exist\n"), input+".xml")
		return
	}

	fmt.Println("File content:\n" + string(file))
}

func deleteXml() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose file name: ")
	input, _ := reader.ReadString('\n')
	input = input[0:len(input)-2] + ""

	if err := os.Remove("data/" + input + ".xml"); err != nil {
		fmt.Printf(Red("File %s does not exist\n"), input+".xml")
		return
	}

	fmt.Println(Green("File deleted"))
}
