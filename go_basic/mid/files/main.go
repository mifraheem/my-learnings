package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to Files handling in Golang")
	data := readFile("./test_file.txt")
	os.Remove("test_file.txt")
	fmt.Println(data, "ifraheem")

	// delete the file

	// fmt.Println(data)
	// content := "I am ifraheem, DevOps engineer."
	// file, err := os.Create("./test_file.txt")

	// defer file.Close()
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println(file)
	// 	leng, err := io.WriteString(file, content)
	// 	if err != nil {
	// 		panic(err)
	// 	} else {
	// 		fmt.Println("File created and written successfully", leng)
	// 	}
	// // }
	// my_data := reafFile(file)
	// fmt.Println(my_data)
}

func readFile(filename string) string {
	data, err := os.ReadFile(filename)
	checkNilErr(err)
	return string(data)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
