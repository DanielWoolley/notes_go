package main

import (
	"fmt"
	"os"
)

func main() {
	Write()
	Read()
}

func Write() {
	file, err := os.Create("notes.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString("notes")
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
