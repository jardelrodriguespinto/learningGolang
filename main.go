package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(addTwoNums(10, 50))
	myStr := "teste"

	var err error = errors.New("Deu ruim")

	if myStr != "teste" {
		fmt.Println(utf8.RuneCountInString(myStr))
	} else {
		fmt.Println(err)
	}
	playingWithMaps()

}

func playingWithArrays() {
	intArr := [3]int{1, 2, 3}
	for i := 0; i < len(intArr); i++ {
		fmt.Println(intArr[i])
	}
}

func addTwoNums(x, y int) int {
	return x + y
}

func exploingTypes() int {
	var intNum int
	intNum = 10
	return intNum
}

type gasEngine struct {
	mpg     int
	gallons int
}

func playingWithSlices() {
	intSlice := [...]int{45, 22, 13, 0}

	for i := 0; i < len(intSlice); i++ {
		if intSlice[i]%2 == 0 {
			intSlice[i] = intSlice[i] + 10
			fmt.Println(intSlice[i])
		}
	}
}

func playingWithMaps() {
	person := map[string]int{"John": 35, "Paul": 42}

	for key, value := range person {
		fmt.Printf("%v is %v years old \n", key, value)
	}
}
