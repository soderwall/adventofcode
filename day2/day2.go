package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	myslice := parseFile("day2.txt")
	var sum int = 0
	for _, row := range myslice {
		//var last int = 0
		for i, _ := range row {
			for y := 0; y < len(row); y++ {
				if i == y {
					// dont do anything
				} else {
					if (row[i] % row[y]) == 0 {
						sum += row[i] / row[y]
					}
				}
			}
		}
	}
	println(sum)
}
func parseFile(filePath string) [][]int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var myslice [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		slice := strings.Split(scanner.Text(), "\t")
		sliceOfInts := make([]int, len(slice))
		for i, row := range slice {
			res, r := strconv.Atoi(row)
			if r != nil {
				log.Fatal(r)
			}
			sliceOfInts[i] = res
		}
		myslice = append(myslice, sliceOfInts)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return myslice
}
