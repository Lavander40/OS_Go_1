package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func showZipMenu() {
	fmt.Println(Yellow("1) Show menu options"))
	//fmt.Println(Yellow("2) Create zip"))
	fmt.Println(Yellow("3) Write in zip"))
	fmt.Println(Yellow("4) Read zip"))
	fmt.Println(Yellow("5) Delete zip"))
	fmt.Println(Yellow("6) Get back"))
}

func zipMenu() {
	showZipMenu()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Your choice: ")
		input, _ := reader.ReadString('\n')
		input = input[0:len(input)-2] + ""
		switch input {
		case "1":
			showZipMenu()
		//case "2":
		//	createZip()
		case "3":
			writeZip()
		case "4":
			readZip()
		case "5":
			deleteZip()
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

func createZip() {
	input := "data/" + getFile() + ".zip"

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
	fmt.Println(Green("File created"))
}

func writeZip() {
	// Сжатие данных
	var buff bytes.Buffer
	reader := bufio.NewReader(os.Stdin)
	zipW := zip.NewWriter(&buff)

	fmt.Print("Choose file to archive (with file format): ")
	input, err := reader.ReadString('\n')
	input = input[0:len(input)-2] + ""
	f, err := zipW.Create(input)
	if err != nil {
		fmt.Print(Red("Invalid input"))
		return
	}
	file, err := os.ReadFile("data/" + input)
	//fmt.Println(Teal(file))
	_, err = f.Write([]byte(file))
	if err != nil {
		fmt.Print(Red("Invalid input"))
		return
	}
	err = zipW.Close()
	if err != nil {
		fmt.Print(Red("Invalid input"))
		return
	}

	// Запись данных в архив
	err = ioutil.WriteFile("data/"+input+".zip", buff.Bytes(), os.ModePerm)
	if err != nil {
		fmt.Print(Red("Invalid input"))
		return
	}

	fmt.Println(Green("File war archived"))
}

func readZip() {
	input := "data/" + getFile() + ".zip"

	zipR, err := zip.OpenReader(input)
	if err != nil {
		fmt.Print(Red("Invalid input"))
		return
	}

	for _, file := range zipR.File {
		fmt.Println(Yellow("Файл " + file.Name + ".zip содержит следующее:"))
		r, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(os.Stdout, r)
		if err != nil {
			fmt.Print(Red("Invalid input"))
			return
		}
		err = r.Close()
		if err != nil {
			fmt.Print(Red("Invalid input"))
			return
		}
	}
}

func deleteZip() {
	input := "data/" + getFile() + ".zip"

	if err := os.Remove(input); err != nil {
		fmt.Print(Red(err))
		return
	}

	fmt.Println(Green("File deleted"))
}
