/*package main

import (
	"fmt"
)

func main() {
	var i int

	fmt.Print("Type a number: ")
	fmt.Scan(&i)
	fmt.Println("Your number is:", i)
}*/

package main

import (
	"fmt"
	"strconv"
)

/*
func main() {
	myslice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
}
*/

type stdent struct {
	firstName  string
	lastName   string
	department string
	address    string
	age        int
}

func main() {
	stdent1 := stdent{firstName: "Asimro", lastName: "Bitaye", department: "Mechanical Engineering", address: "Addis Ababa", age: 27}
	Stdent2 := stdent{firstName: "Tlahun", lastName: "Yirga", department: "civil Engineering", address: "Addis Ababa", age: 27}

	fmt.Println("hello, my name is " + stdent1.firstName + " " + stdent1.lastName + ". i am graduated in " + stdent1.department + " right now i am in " + stdent1.address + " and i am " + strconv.Itoa(stdent1.age) + " years old.")
	Stdent2.age++
	fmt.Println("hello, my name is " + Stdent2.firstName + " " + Stdent2.lastName + ". i am graduated in " + Stdent2.department + " right now i am in " + Stdent2.address + " and i am " + strconv.Itoa(Stdent2.age) + " years old.")
}
