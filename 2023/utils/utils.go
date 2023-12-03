package utils

import (
	"bufio"
	"os"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

type Parser func(string, int) []interface{}

func ParseInput(lineParser Parser, filePath string) []interface{} {
	file, err := os.Open(filePath)
	CheckError(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var parsedItems []interface{}
	currentLine := 0

	for scanner.Scan() {
		line := scanner.Text()

		parsedLine := lineParser(line, currentLine)
		currentLine++

		if parsedLine != nil {
			parsedItems = append(parsedItems, parsedLine...)
		}
	}

	return parsedItems
}
