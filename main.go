package main

import (
	"fmt"
	"learn-go/utils"
)

func showName(name string) string {
	return name
}

func sum(a int, b int) int {
	return a + b
}

func multipleSum(a int, b int, c int, d int) (firstSumParams int, secondSumParams int) {
	firstSum := a + b
	secondSum := c + d
	return firstSum, secondSum
}

func conditionF(name string) string {
	if name == "Mazharul" {
		return " This is Mazharul"
	} else if name == "arif" {
		return " This is Arif"
	} else {
		return "Not Mazharul"
	}
}

func main() {
	name := "Mazharul"
	age := 25
	height := 5.6
	isProgrammer := true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Is Programmer:", isProgrammer)
	myName := showName("Md Mazharul Islam")
	fmt.Println(myName)

	mySum := sum(10, 20)
	fmt.Println(mySum)
	firstSum, secondSum := multipleSum(10, 20, 30, 40)
	fmt.Println(firstSum, secondSum)
	newSum := utils.Sum(5, 5)
	fmt.Println(newSum)
	nameMe, ageMe := utils.Introduce("Md Mazharul Islam", 25)
	fmt.Println("My name is", nameMe, "and I am", ageMe, "years old.")
	findMazharul := conditionF("arif")
	fmt.Println(findMazharul)

}
