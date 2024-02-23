package main

import (
	"fmt"
	"os"
)

func main() {
	var date string
	fmt.Print("Enter the date")
	fmt.Scan(&date)
	Write(date)
	Read()
}

func Write(date string) {
	file, err := os.Create("notes.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(date)
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
