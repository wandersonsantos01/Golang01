package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const tries = 3
const sleepSeconds = 5

// import "reflect"

func main() {
	name, age := getNameAndAge()
	_, ageTwo := getNameAndAge()

	fmt.Println(ageTwo)

	showIntroduction(name, age)

	for {
		showMenu()
		command := readCommand()

		// if command == 1 {
		// 	fmt.Println("Monitoring")
		// } else if command == 2 {
		// 	fmt.Println("Showing logs")
		// } else if command == 0 {
		// 	fmt.Println("Bye...")
		// } else {
		// 	fmt.Println("Unknown command")
		// }

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing logs")
		case 0:
			fmt.Println("Bye...")
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
			os.Exit(-1)
		}
		fmt.Println("")
	}
}

func showIntroduction(name string, age int) {
	// var name string = "John"
	// var version float32 = 1.3
	// var age = 34

	// name := "John"
	version := 1.3
	// age := 34

	fmt.Println("Hello Mr.", name, ", your age is", age)
	fmt.Println("Versão", version)

	// fmt.Println("Type of variable is", reflect.TypeOf(name))
	// fmt.Println("Type of variable is", reflect.TypeOf(version))
	// fmt.Println("Type of variable is", reflect.TypeOf(age))
}

func showMenu() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")
}

func readCommand() int {
	var command int
	fmt.Scanf("%d", &command)
	return command
}

func getNameAndAge() (string, int) {
	name := "John"
	age := 34
	return name, age
}

func startMonitoring() {
	fmt.Println("Monitoring")

	// arrays
	// var sitesArray [4]string
	// sitesArray[0] = "https://www.alura.com.br"
	// sitesArray[1] = "https://httpstat.us/Random/200,400"
	// sitesArray[2] = "https://caelum.com.br"
	// sitesArray[3] = "https://wandersonsantos01.github.io"

	// slices
	// sitesSlice := []string{"https://www.alura.com.br", "https://httpstat.us/Random/200,400", "https://caelum.com.br", "https://wandersonsantos01.github.io"}

	sitesSlice := readSitesFile()

	for i := 0; i < tries; i++ {
		for idx, site := range sitesSlice {
			fmt.Println("Testing site", idx, ":", site)
			checkSite(site)
		}
		fmt.Println("")
		time.Sleep(sleepSeconds * time.Second)
	}

}

func checkSite(site string) {
	res, _ := http.Get(site)

	if res.StatusCode == 200 {
		fmt.Println("Site:", site, "OK")
	} else {
		fmt.Println("Site:", site, "ERROR - Status Code:", res.StatusCode)
	}
	fmt.Println("=========================================================")
}

func readSitesFile() []string{
	file, _ = os.Open("sites.txt")
}