package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// remove remove item from strings input by index
func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}

// findSmallest find smallest index in strings input
func findSmallest(data []string) int {
	smallest := data[0]
	idx := 0

	for i := 1; i < len(data); i++ {
		if data[i] < smallest {
			smallest = data[i]
			idx = i
		}
	}

	return idx
}

// selectionSort sort input of strings
func selectionSort(data []string) []string {
	var slice []string

	for range data {
		smallest := findSmallest(data)
		slice = append(slice, data[smallest])
		data = remove(data, smallest)
	}

	return slice
}

// selectionSort sort input of strings (another selectionSort func)
//func selectionSort(data []string) []string {
//	for i := 0; i < len(data); i++ {
//		var minIdx = i
//		for j := i; j < len(data); j++ {
//			if data[j] < data[minIdx] {
//				minIdx = j
//			}
//		}
//		data[i], data[minIdx] = data[minIdx], data[i]
//	}
//
//	return data
//}

// run start program
func run(filepath string) {
	data, err := readFile(filepath)
	if err != nil {
		log.Fatalf("can't read file: %v", err)
	}

	fmt.Println(selectionSort(data))
}

// readFile read file by path into input of strings
func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	sc := bufio.NewScanner(file)
	var slice []string
	for sc.Scan() {
		slice = append(slice, sc.Text())
	}

	return slice, sc.Err()
}

func main() {
	if len(os.Args) == 2 {
		run(os.Args[1])
		return
	}
	log.Fatalln("add file as arg")
}
