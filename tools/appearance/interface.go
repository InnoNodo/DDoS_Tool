package appearance

import (
	"fmt"
	"log"
	"os"
)

var URL string
var Reset = "\033[0m"
var Red = "\033[31;1m"
var Green = "\033[32;1m"
var Yellow = "\033[33;1m"

func Banner() string {
	file, err := os.ReadFile("tools/appearance/banner.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Red + string(file) + Reset + "\n")
	fmt.Println(Yellow + "Creator: Nodo")
	fmt.Print(Yellow + "Type URL/IP: " + Reset)
	fmt.Scanln(&URL)
	return URL
}
