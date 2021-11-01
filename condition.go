package main

import (
	"fmt"
	"runtime"
)

func main() {
	//name1 := "nana"
	//
	//if name1 == "Nic"{
	//	fmt.Println("hello " + name1)
	//} else {
	//	fmt.Println("You is " + name1)
	//}
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "windows":
		fmt.Println("Windows.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
