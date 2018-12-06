package main

import (
	"log"
	"os"
	"strconv"
	"github.com/shreshthkhilani/aoc18/reader"
)

func main() {
	filePath := os.Args[1]
	if filePath == "" {
		log.Fatal("No input file path specified.")
	}
	
	lines, err := reader.GetLines(filePath)
	if err != nil {
		log.Fatal(err)
	}
	
	data := LineParser(lines)
	log.Println(Part1(data, 0))
	log.Println(Part2(data, make(map[int]bool), 0))
}

func LineParser(lines []string) []int {
	var data []int
	for _, line := range lines {
		datum, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			log.Println("Skipping line:", line)
			continue
		}
		data = append(data, int(datum))
	}
	return data
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