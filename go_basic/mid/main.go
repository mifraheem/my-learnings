package main

import "fmt"

func main() {
	// var languages = make(map[string]string)
	// languages["Py"] = "python"
	// languages["Py"] = "python"
	// languages["Js"] = "javascript"
	// languages["Js"] = "javascript"
	// fmt.Println("The langugaes are: ", languages)
	// var varpy string = languages["Py"]
	// if varpy == "" {
	// 	fmt.Println("Data not found ")
	// } else {
	// 	fmt.Println("The value of Py is: ", varpy)
	// }

	// delete(languages, "Py")

	// for key, value := range languages {
	// 	fmt.Println(key, value)
	// }

	// var number int = 40
	// switch number {
	// case 50:
	// 	fmt.Println("i got number 50")

	// case 40:
	// 	fmt.Println("i got 40 > ")

	// default:
	// 	fmt.Println("this is in default")
	// 	fmt.Println("the final line is great")
	// }
	// rand.Seed(time.Now().UnixNano())
	// dice_number := rand.Intn(6) + 1
	// fmt.Println("you rolled a", dice_number)
	// days := []string{"Sunday", "Monday", "Tuesday", "Wed", "Thur", "Fri", "Sat"}
	// for _, day := range days {
	// 	fmt.Println(day)
	// }
	// var data = proAdder(3, 4, 5, 2, 4, 5)
	// fmt.Println(data)

	nme := "Hurair"
	rll := 20
	name, roll_number := nameAndNumber(nme, rll)
	fmt.Println(name, roll_number)

}

func proAdder(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func nameAndNumber(name string, roll int) (int, string) {
	return roll, name
}
