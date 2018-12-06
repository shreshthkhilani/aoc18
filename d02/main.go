package main

import (
	"log"
	"os"
	"strings"
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
	
	data := lines
	log.Println(Part1(data))
	log.Println(Part2(data))
}

func Part1(data []string) int {
	twos := 0
	threes := 0
	for _, datum := range data {
		hasTwo := false
		hasThree := false
		for _, c := range datum {
			count := strings.Count(datum, string(c))
			if count == 2 {
				hasTwo = true;
			} else if count == 3 {
				hasThree = true;
			}
		}
		if hasTwo {
			twos += 1
		}
		if hasThree {
			threes += 1
		}
	}
	return twos * threes
}

func DiffersByOne(a string, b string) (bool, string) {
	if len(a) != len(b) {
		return false, ""
	}
	var runesA, runesB = []rune(a), []rune(b)
	var runesCommon []rune
	for i := 0; i < len(runesA); i++ {
		if runesA[i] == runesB[i] {
			runesCommon = append(runesCommon, runesA[i])
		}
	}
	if len(runesCommon) + 1 == len(runesA) {
		return true, string(runesCommon)
	}
	return false, ""
}

func Part2(data []string) string {
	for i, a := range data {
		for _, b := range data[i + 1:] {
			doesDifferByOne, commonString := DiffersByOne(a, b)
			if doesDifferByOne {
				return commonString
			}
		}
	}
	return ""
}
