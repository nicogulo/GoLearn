package main

import "fmt"

//type Vertex struct {
//	X, Y int
//}

//var (
//	v1 = Vertex{2, 4}  // has type Vertex
//	v2 = Vertex{X: 1}  // Y:0 is implicit
//	v3 = Vertex{}      // X:0 and Y:0
//	p  = &Vertex{1, 2} // has type *Vertex
//)

func main() {

	// fmt.Println(v1, p, v2, v3)

	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	// fmt.Println("Hello World")
	// fmt.Println("satu = ", 1)
	// fmt.Println("Dua = ", 2)
	// fmt.Println("Tiga koma satu = ", 3.1)
	// fmt.Println(len("Nic"))

	//create multiple variable
	//var (
	//	firstName = "Dev"
	//	lastName = "Nic"
	//)
	//fmt.Println(firstName)
	//fmt.Println(lastName)
	//
	//const testConst = "test contstant"
	//fmt.Println(testConst)

	//Conversion string
	//var (
	//	name = "Nic"
	//	e byte = name[0]
	//	eString string = string(e)
	//)
	//fmt.Println(eString)

	//math operation
	//const (
	//	numba = 12
	//	numbb = 3
	//	i =+ numba + numbb
	//)
	//fmt.Println(i)

	//array
	//var name [3]string
	//name[0] = "satu"
	//name[1] = "dua"
	//name[2] = "tiga"
	//
	//fmt.Println(name[2])
	//v := true // change me!
	//fmt.Printf("v is of type %T\n", v)

}
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
