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
	black   = Color("\033[1;30m%s\033[0m")
	red     = Color("\033[1;31m%s\033[0m")
	green   = Color("\033[1;32m%s\033[0m")
	yellow  = Color("\033[1;33m%s\033[0m")
	purple  = Color("\033[1;34m%s\033[0m")
	magenta = Color("\033[1;35m%s\033[0m")
	teal    = Color("\033[1;36m%s\033[0m")
	white   = Color("\033[1;37m%s\033[0m")
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
			os.Exit(0)
		case "4":
			logicDisks()
		case "5":
			fileMenu()
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
	fmt.Println(yellow("3) Exit"))
	fmt.Println(yellow("4) Show disk info"))
	fmt.Println(yellow("5) Show file menu"))
}

func showFileMenu() {
	fmt.Println(yellow("1) Show menu options"))
	fmt.Println(yellow("2) Create file"))
	fmt.Println(yellow("3) Read file"))
	fmt.Println(yellow("4) Delete file"))
	fmt.Println(yellow("5) Get back"))
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
			readFile()
		case "4":
			deleteFile()
		case "5":
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

	file, err := os.Create(input)

	if err != nil {
		log.Panic(err)
	}

	defer file.Close()
	fmt.Println("File created")
}

func readFile() {
	fmt.Println("File read")
}

func deleteFile() {
	fmt.Println("File deleted")
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
		log.Panic(err)
	}
	fmt.Println(magenta(string(stdout)))
}

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}
