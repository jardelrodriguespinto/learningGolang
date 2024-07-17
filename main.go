package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {

}

func leaningAboutChannels() {
	channel := make(chan int)
	go setList(channel)

	//	valor := <- channel

	for v := range channel {
		fmt.Println("recebendo: ", v)
	}

}

func setList(channel chan<- int) {
	for i := 0; i < 100; i++ {
		channel <- i
		fmt.Println("enviando: ", i)
	}
	close(channel)

}

func mutext() {
	var m sync.Mutex

	i := 0
	for x := 0; x < 10000; x++ {
		go func() {
			m.Lock()
			i++
			m.Unlock()
		}()

	}

	time.Sleep(time.Second * 2)
	fmt.Println(i)

}

func ChangeNumber(i *int, newNumber int) {
	*i = newNumber
}

func awaitGroup() {
	var wg sync.WaitGroup
	wg.Add(3)
	go CallDatabase(&wg)
	go CallApi(&wg)
	go ProcessInternal(&wg)

	wg.Wait()

}

func gorotines() {
	for i := 0; i < 100; i++ {
		go showMessage(strconv.Itoa(i))
	}

	time.Sleep(time.Millisecond * 2)

}

func CallDatabase(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("Calling the database")
	wg.Done()
}

func CallApi(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("Calling the api")
	wg.Done()
}

func ProcessInternal(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("Calling the internal process")
	wg.Done()
}

func showMessage(message string) {
	fmt.Println(message)
}

func playingWithGenerics() {
	arrInt := []int{1, 2, 3, 4, 5, 6}
	arrString := []string{"a", "b", "c", "d", "e", "f"}

	fmt.Println(reverseArr(arrInt))
	fmt.Println(reverseArr(arrString))
}

type constaint interface {
	int | string
}

func binarySearch(arr []int, target int) int {

	start := arr[0]
	end := len(arr) - 1

	for i := 0; i < len(arr); i++ {
		mid := start + end/2

		if arr[mid] == target {
			return i
		} else if arr[mid] < target {
			mid = start + 1
		} else {
			mid = end - 1
		}

	}

	return -1
}

func reverseArr[T constaint](slice []T) []T {

	newArr := make([]T, len(slice))

	tamanhoNewArr := 0

	for i := len(slice) - 1; i >= 0; i-- {
		if tamanhoNewArr < len(newArr) {
			newArr[i] = slice[tamanhoNewArr]
			tamanhoNewArr++
		}
	}

	return newArr
}
