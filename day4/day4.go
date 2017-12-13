package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data := parseFile("day4.txt")
	part1(data)
	part2(data)
}

func part2(passphrases [][]string) {
	var validPassphrases int
	for _, row := range passphrases {
		var valid = true
		used := make(map[string]bool)
		for _, word := range row {
			for k := range used {
				keySplit := strings.Split(k, "")
				sort.Strings(keySplit)
				wordSplit := strings.Split(word, "")
				sort.Strings(wordSplit)
				if reflect.DeepEqual(keySplit, wordSplit) {
					valid = false
					break
				}
			}
			if valid {
				used[word] = true
			} else {
				break
			}
		}
		if valid {
			validPassphrases++
		}
	}
	fmt.Println(validPassphrases)
}
func part1(passphrases [][]string) {
	var validPassphrases int
	var unique bool
	var tmpSlice []string
	for _, passphrase := range passphrases {
		unique = true
		for i, word := range passphrase {
			//fmt.Print("passphrase:")
			//fmt.Println(word)
			tmpSlice = removdIndex(passphrase, i)
			if findWord(tmpSlice, word) == false {
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
	f, err := os.Open(filePath)
	check(err)
	defer f.Close()

	// Parse file
	data := make([][]string, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		row := strings.Fields(s.Text())
		data = append(data, row)
	}
	return data
}
