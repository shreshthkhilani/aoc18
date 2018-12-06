package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filePath := os.Args[1]
	if filePath == "":
		fmt.Println("No input file path specified.")
	dat, err := ioutil.ReadFile(filePath)
	if err != nil:
		fmt.Println("Unable to read file at path:", filePath)
		os.Exit(1)
	fmt.Print(string(dat))
	
	var data []int
	data = append(data, 3, 3, 4, -2, -4)
	fmt.Println(Part1(data, 0))
	fmt.Println(Part2(data, make(map[int]bool), 0))
}

func Part1(data []int, frequency int) int {
	for _, v := range data {
		frequency += v
	}
	return frequency
}

func Part2(data []int, m map[int]bool, frequency int) int {
	for _, v := range data {
		frequency += v
		_, ok := m[frequency]
		if ok {
			return frequency
		}
		m[frequency] = true
	}
	return Part2(data, m, frequency)
}