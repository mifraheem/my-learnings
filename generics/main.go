package main

import "fmt"

func printSlice[T any](slice []T) {
	for _, sl := range slice {
		fmt.Println(sl)
	}
}

func main() {
	var s1 []int = []int{1, 2, 3, 4, 5}
	subjects := []string{"math", "urdu", "english"}
	printSlice(s1)
	printSlice(subjects)

}
