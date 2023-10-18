package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	Name
	Age int
}

type Name struct {
	FirstName string
	LastName  string
}

func NewUser(firstName, lastName string) *User {
	return &User{
		Name: Name{
			FirstName: firstName,
			LastName:  lastName,
		},
		Age: 15,
	}
}

func (u User) Serialize() string {
	return fmt.Sprintf("Serialize method called\nname: %s %s\n", u.FirstName, u.FirstName)
}

func showJsonMenu() {
	fmt.Println(Yellow("1) Show menu options"))
	fmt.Println(Yellow("2) Create json"))
	fmt.Println(Yellow("3) Write in json"))
	fmt.Println(Yellow("4) Read json"))
	fmt.Println(Yellow("5) Delete json"))
	fmt.Println(Yellow("6) Get back"))
}

func jsonMenu() {
	showJsonMenu()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Your choice: ")
		input, _ := reader.ReadString('\n')
		input = input[0:len(input)-2] + ""
		switch input {
		case "1":
			showJsonMenu()
		case "2":
			createJson()
		case "3":
			writeJson()
		case "4":
			readJson()
		case "5":
			deleteJson()
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

func createJson() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose file name: ")
	input, _ := reader.ReadString('\n')
	input = input[0:len(input)-2] + ""

	_, err := os.Stat("data/" + input + ".json")

	if err == nil {
		fmt.Println(Teal("File " + input + " already exist, rewrite? (print y)"))
		answer, _ := reader.ReadString('\n')
		answer = answer[0:len(answer)-2] + ""
		if answer != "y" {
			return
		}
	}

	file, err := os.Create("data/" + input + ".json")

	if err != nil {
		fmt.Print(err)
		return
	}

	defer file.Close()
	fmt.Println(Green("File created"))
}

func writeJson() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose file name: ")
	input, _ := reader.ReadString('\n')
	input = input[0:len(input)-2] + ""

	file, err := os.Create("data/" + input + ".json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("File was opened\nCreating User instance...")
	fmt.Println(Yellow("Input first and last names:"))
	firstName, _ := reader.ReadString('\n')
	firstName = firstName[0:len(firstName)-2] + ""
	lastName, _ := reader.ReadString('\n')
	lastName = lastName[0:len(lastName)-2] + ""

	user := NewUser(firstName, lastName)
	content, err := json.Marshal(user)

	_, err = file.WriteString(string(content))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Green("Writing complete"))
}

func readJson() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose file name: ")
	input, _ := reader.ReadString('\n')
	input = input[0:len(input)-2] + ""

	//user := User{}
	file, err := os.ReadFile("data/" + input + ".json")
	//err = json.Unmarshal(file, &user)
	if err != nil {
		fmt.Print(Red(err))
		return
	}

	fmt.Println("File content:\n" + string(file)) //+ user.Serialize())
}

func deleteJson() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose file name: ")
	input, _ := reader.ReadString('\n')
	input = input[0:len(input)-2] + ""

	if err := os.Remove("data/" + input + ".json"); err != nil {
		fmt.Print(Red(err))
		return
	}

	fmt.Println(Green("File deleted"))
}
