package main

import "fmt"

/*
const (
	a int = 1
	b     = 3.14
	c     = "hello"
	d     = true
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}*/

func main() {
	/*var i, j string = "hello", "world"

	print(i, " ", j)*/

	/*var a, b = 10, 20

	print(a, b)*/

	/*var i string = "hello"
	var j int = 15

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T\n", j, j)*/

	/*const premiumPlanName = "premium plan"
	const basicPlanName = "Basic plan"

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)*/

	/*const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	fmt.Println("one hour is equal to", secondsInHour, "seconds")*/

	const name = "kim john"
	const openRate = 30.546

	msg := fmt.Sprint("hi %S, your open rate is: %.1f persent.",
		name,
		openRate)

	fmt.Println(msg)

}
