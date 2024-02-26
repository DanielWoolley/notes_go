package main

import (
	"fmt"
	"os"
)

func main() {
	date, note := getUserInput()
	Write(date, note)
	Read()
}

func getUserInput() (string, string) {
	var date string
	var note string
	fmt.Print("Enter the date ")
	fmt.Scan(&date)
	fmt.Print("Enter your notes ")
	fmt.Scan(&note)
	return date, note
}

func Write(date string, note string) {
	file, err := os.Create("notes.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(date + "\n" + note)
	if err != nil {
		panic(err)
	}
}

func Read() {
	data, err := os.ReadFile("notes.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
