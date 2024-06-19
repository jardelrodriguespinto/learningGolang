package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	arr := []int{20, 5, 1, 5, 2}
	bubbleSort(arr)
	mapsInGo()
	iterandoSobreMap()
}

func testandoArrays() {

	arr := []int{54, 84, 55}

	novaLista := arr[:3]

	fmt.Println(novaLista)

}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}

	fmt.Println(arr)
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

func exemplo() (int, error) {
	numero := 15
	erro := errors.New("teste")
	return numero, erro
}

func playingWithVariables() (int, reflect.Type) {
	var numero int = 32
	numero2 := 15
	return numero + numero2, reflect.TypeOf(numero)
}

func playingWithConstants() string {
	const text string = "teste"
	//text = "10"
	return text
}

func playingWithWhileAndForLoop() {
	arr := []int{1, 2, 3, 80}

	i := 0
	tamanho := len(arr)

	for i < tamanho {
		fmt.Println(arr[i])
		i++
	}

}

func mapsInGo() {
	myMap := make(map[string]int)
	myList := make([]int, 1)
	myList = append(myList, 12)

	myMap["Jo"] = 10
	myMap["as"] = 10

	valor, ok := myMap["teste"]

	if ok {
		fmt.Println(valor)
	} else {
		fmt.Println("valor não existe")
	}

	for i := 0; i < len(myList); i++ {
		fmt.Println(myList[i])
	}

}

func iterandoSobreMap() {
	pessoa := make(map[string]int)

	pessoa["John"] = 45
	pessoa["Mary"] = 16

	for chave, valor := range pessoa {
		fmt.Println(chave + " tem " + strconv.Itoa(valor) + " anos de idade")
	}

}
