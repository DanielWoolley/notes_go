package main

import (
	"fmt"
	"os"
)

func main() {
	date := getDate()
	note := getNote()
	Write(date, note)
	Read()
}

func getDate() string {
	var date string
	fmt.Println("Enter the date ")
	fmt.Scan(&date)
	return date
}

func getNote() string {
	var note string
	fmt.Print("Enter your notes ")
	fmt.Scan(&note)
	return note
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
