package main

import "os"

func main() {
	Write()
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
