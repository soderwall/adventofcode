package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	passphrases := parseFile("day4	.txt")
	var validPassphrases int
	var unique bool
	var tmpSlice []string
	for _, passphrase := range passphrases {
		unique = true
		fmt.Println("++++++++++++++++++++")
		for i, word := range passphrase {
			//fmt.Print("passphrase:")
			//fmt.Println(word)
			tmpSlice = removdIndex(passphrase, i)
			if findWord(tmpSlice, word) == false {
				fmt.Println("FOUND MATCH FOR: ")
				fmt.Print(word)
				fmt.Println(tmpSlice)
				unique = false
			}
		}
		if unique == true {
			validPassphrases++
		}

	}
	fmt.Println(validPassphrases)
}
func removdIndex(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
func findWord(slice []string, word string) bool {
	for _, wordToMatch := range slice {
		//cba to break this out into the parseFile. I suspect the scanner API returns some
		// strange characters. :/
		wordToMatch = strings.ToLower(strings.Trim(wordToMatch, " \r\n"))
		word = strings.ToLower(strings.Trim(word, " \r\n"))
		if strings.Compare(wordToMatch, word) == 0 {
			return false
		}
	}
	return true
}
func parseFile(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var myslice [][]string
	var slice []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		slice = strings.SplitAfter(scanner.Text(), " ")
		myslice = append(myslice, slice)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return myslice
}
