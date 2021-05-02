package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
)

// - preprocessing input
func linesFromReader(r io.Reader) (string, error) {
	var input string
	reader := bufio.NewReader(r)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		input += fmt.Sprintf("%s", line)
	}

	return input, nil
}

func inputPreprocessing(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return linesFromReader(resp.Body)
}

func removeDuplication(input string) string {
	result := ""
	mapper := make(map[string]bool)
	slice := []string{}

	for _, val := range input {
		slice = append(slice, fmt.Sprintf("%c", val))
	}

	for _, valSlice := range slice {
		if mapper[valSlice] == false {
			mapper[valSlice] = true
			result += valSlice
		}
	}

	return result
}

// Insert a char in a word
func insertAt(i int, char string, perm string) string {
	start := perm[0:i]
	end := perm[i:len(perm)]
	return start + char + end
}

func lexicographic(input string) []string {
	// - base case, for one char, all perms are [char]
	if len(input) == 1 {
		return []string{input}
	}

	current := input[0:1]
	remStr := input[1:]

	perms := lexicographic(remStr)

	result := make([]string, 0)

	// - for every value in the perms of substring
	for _, val := range perms {
		// - add current char at every possible position
		for x := 0; x <= len(val); x++ {
			newPerm := insertAt(x, current, val)
			fmt.Println("Str : ", newPerm)
			result = append(result, newPerm)
		}
	}

	return result
}

func main() {
	// - raw data
	url := "https://gist.githubusercontent.com/Jekiwijaya/0b85de3b9ff551a879896dd78256e9b8/raw/e9d58da5d4df913ad62e6e8dd83c936090ee6ef4/gistfile1.txt"

	// - preprocessing data from URL
	prepResult, err := inputPreprocessing(url)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	// - remove duplication from string;
	uniqueValue := removeDuplication(prepResult)
	fmt.Println("UniqueValue : ", uniqueValue)

	// - logic for lexicographic;
	result := lexicographic(uniqueValue)
	fmt.Println("Result : ", result)
}
