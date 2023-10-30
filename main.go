package main

import (
	"bufio"
	"fmt"
	"github.com/shirou/gopsutil/host"
	"os"
	"runtime"
	"strconv"
)

var OS = runtime.GOOS
var reader = bufio.NewReader(os.Stdin)

func showMenu() {
	fmt.Println(Yellow("1) Show menu options"))
	fmt.Println(Yellow("2) Show OS type"))
	fmt.Println(Yellow("3) Show disk info"))
	fmt.Println(Yellow("4) Show file menu"))
	fmt.Println(Yellow("5) Show json menu"))
	fmt.Println(Yellow("6) Show xml menu"))
	fmt.Println(Yellow("7) Show zip menu"))
	fmt.Println(Yellow("8) Exit"))
}

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
			jsonMenu()
		case "6":
			xmlMenu()
		case "7":
			zipMenu()
		case "8":
			os.Exit(0)
		default:
			fmt.Println(Red("Invalid choice."))
			continue
		}
		//fmt.Println("You chose option " + Yellow((input))
	}
}

func OSType() {
	fmt.Println("OS: " + Yellow(OS))
	hostStat, _ := host.Info()
	fmt.Println("Version: " + Yellow(hostStat.Platform))
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
		fmt.Println("Disk " + strconv.Itoa(i) + " : " + Yellow(val))
	}
}
