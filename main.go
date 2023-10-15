package main

import (
	"bufio"
	"fmt"
	"github.com/shirou/gopsutil/host"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

var (
	//black   = Color("\033[1;30m%s\033[0m")
	//white   = Color("\033[1;37m%s\033[0m")
	//purple  = Color("\033[1;34m%s\033[0m")
	//magenta = Color("\033[1;35m%s\033[0m")
	red    = Color("\033[1;31m%s\033[0m")
	green  = Color("\033[1;32m%s\033[0m")
	yellow = Color("\033[1;33m%s\033[0m")
	teal   = Color("\033[1;36m%s\033[0m")
)

var OS = runtime.GOOS

func main() {
	showMenu()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Your choice: ")
		input, _ := reader.ReadString('\n')
		input = input[0:len(input)-2] + ""
		switch input {
		case "1":
			showMenu()
		case "2":
			OSType()
		case "3":
			logicDisks()
		case "4":
			fileMenu()
		case "5":
			os.Exit(0)
		default:
			fmt.Println(red("Invalid choice."))
			continue
		}
		//fmt.Println("You chose option " + yellow(input))
	}
}

func showMenu() {
	fmt.Println(yellow("1) Show menu options"))
	fmt.Println(yellow("2) Show OS type"))
	fmt.Println(yellow("3) Show disk info"))
	fmt.Println(yellow("4) Show file menu"))
	fmt.Println(yellow("5) Exit"))
}

func showFileMenu() {
	fmt.Println(yellow("1) Show menu options"))
	fmt.Println(yellow("2) Create file"))
	fmt.Println(yellow("3) Write in file"))
	fmt.Println(yellow("4) Read file"))
	fmt.Println(yellow("5) Delete file"))
	fmt.Println(yellow("6) Get back"))
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
			fmt.Println(red("Invalid choice."))
			continue
		}
		//fmt.Println("You chose option " + yellow(input))
	}
}

func createFile() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose file name: ")
	input, _ := reader.ReadString('\n')
	input = input[0:len(input)-2] + ""

	_, err := os.Stat(input)

	if err == nil {
		fmt.Println(teal("File " + input + " already exist, rewrite? (print y)"))
		answer, _ := reader.ReadString('\n')
		answer = answer[0:len(answer)-2] + ""
		if answer != "y" {
			return
		}
	}

	file, err := os.Create(input)

	if err != nil {
		fmt.Print(err)
		return
	}

	defer file.Close()
	fmt.Println(green("File created"))
}

func writeFile() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose file name: ")
	input, _ := reader.ReadString('\n')
	input = input[0:len(input)-2] + ""

	file, err := os.Create(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Print(yellow("File was opened\nWrite text to input: "))
	input, _ = reader.ReadString('\n')

	_, err = file.WriteString(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(green("Writing complete"))
}

func readFile() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose file name: ")
	input, _ := reader.ReadString('\n')
	input = input[0:len(input)-2] + ""

	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Print(red(err))
		return
	}

	fmt.Println("File content:\n" + string(file))
}

func deleteFile() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose file name: ")
	input, _ := reader.ReadString('\n')
	input = input[0:len(input)-2] + ""

	if err := os.Remove(input); err != nil {
		fmt.Print(red(err))
		return
	}

	fmt.Println(green("File deleted"))
}

func OSType() {
	fmt.Println("OS: " + yellow(OS))
	hostStat, _ := host.Info()
	fmt.Println("Version: " + yellow(hostStat.Platform))
}

func logicDisks() {
	var r []string
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, err := os.Open(string(drive) + ":\\")
		if err == nil {
			r = append(r, string(drive))
			f.Close()
		}
	}
	for i, val := range r {
		fmt.Println("Disk " + strconv.Itoa(i) + " : " + yellow(val))
	}
}

func testExec() {
	cmd := exec.Command("systeminfo")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(stdout))
}

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}
