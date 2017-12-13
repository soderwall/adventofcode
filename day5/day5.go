package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data := parseFile("day5.txt")
	fmt.Printf("%v", data)
	data2 := make([]int, len(data))
	copy(data2, data)
	part1(data)
	part2(data2)
}

func part2(data []int) {
	fmt.Println("Starting to walk part 2...!")
	var jumps, index = 0, 0
	for {
		jumps++
		currentValue := data[index]
		if currentValue >= 3 {
			data[index]--
		} else {
			data[index]++
		}
		index += currentValue
		if index >= len(data) || index < 0 {
			fmt.Println(jumps)
			return
		}
	}
}
func part1(data []int) {

	fmt.Println("Starting to walk part 1...")
	var jumps, index = 0, 0
	for {
		jumps++
		currentValue := data[index]
		data[index]++
		index += currentValue
		if index >= len(data) || index < 0 {
			fmt.Println(jumps)
			return
		}
	}
}

func parseFile(filePath string) []int {
	f, err := os.Open(filePath)
	check(err)
	defer f.Close()

	// Parse file
	data := make([]int, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		row := s.Text()
		res, _ := strconv.Atoi(row)
		data = append(data, res)
	}
	return data
}
