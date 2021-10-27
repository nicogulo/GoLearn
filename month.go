package main

import "fmt"

func main() {
	var month = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}
	slice1 := month[11:]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	slice1[0] = "meiyaa"
	fmt.Println(slice1)
	fmt.Println(month)

	person := map[string]string{
		"name":    "Nic",
		"address": "Jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// 	make map
	book := make(map[string]string)
	book["title"] = "GoLearn"
	book["author"] = "Backend"

	fmt.Println(book["title"])
	fmt.Println(book["author"])
}
