package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	exerciseCourse2()
}

func testandoArrays() {

	arr := []int{54, 84, 55}

	novaLista := arr[:3]

	fmt.Println(novaLista)

}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}

	return arr
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

func exerciseCourse() {
	myArr := make([]int, 2)
	myArr[0] = 15
	myArr[1] = 10

	count := 0

	for i := 0; i < len(myArr); i++ {
		count += myArr[i]
	}

	fmt.Println(count)
}

func exerciseCourse2() {
	myArr := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}

	newArr := bubbleSort(myArr)

	result := sumAll(newArr[:5]) + sumAll(newArr[5:])

	fmt.Println(result)
}

func sumAll(arr []int) int {
	count := 0

	for i := 0; i < len(arr); i++ {
		count += arr[i]
	}

	return count
}

func calculadora(num1 int, num2 int) (int, int, int, int) {
	soma := num1 + num2
	subtracao := num1 - num2
	divisao := num1 / num2
	multiplicao := num1 * num2

	return soma, subtracao, divisao, multiplicao
}
