package main

import (
	"fmt"
	"os"
)

const fileName = "notes.text"

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
	file, err := os.Create(fileName)
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
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
