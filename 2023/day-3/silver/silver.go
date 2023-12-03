package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"utils"
)

type point struct {
	x int
	y int
}

var maxPoint = point{x: 140, y: 140}

const symbols = "*%#@+-=/&$"

type coordinate struct {
	location point
	isSymbol bool
	value    rune
}

type numberCoordinate struct {
	locations []point
	value     int
}

func isValidPosition(location point) bool {
	x := location.x
	y := location.y

	if x < 0 || y < 0 || x > maxPoint.x || y > maxPoint.y {
		return false
	}

	return true
}

func (location point) getAdjacentCells() []point {
	var adjacentCells []point

	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if x == 0 && y == 0 {
				continue
			}

			adjacentCell := point{x: location.x + x, y: location.y + y}

			if isValidPosition(adjacentCell) {
				adjacentCells = append(adjacentCells, adjacentCell)
			}
		}
	}

	return adjacentCells
}

func getSymbols(line string, lineNumber int) []interface{} {
	var parsedSymbols []interface{}

	for x, value := range line {
		if value == '.' {
			continue
		}

		if strings.ContainsRune(symbols, value) {
			location := point{x: x, y: lineNumber}
			parsedSymbol := coordinate{location: location, isSymbol: true, value: value}

			parsedSymbols = append(parsedSymbols, parsedSymbol)
			continue
		}

	}

	return parsedSymbols
}

func getNumbers(line string, lineNumber int) []interface{} {
	var parsedNumbers []interface{}

	currentNumber := ""
	var currentNumberLocations []point

	index := 0

	for index < len(line) {
		currentNumPoint := point{x: index, y: lineNumber}

		if unicode.IsNumber(rune(line[index])) {
			currentNumber = currentNumber + string(line[index])
			currentNumberLocations = append(currentNumberLocations, currentNumPoint)
			index++
			continue
		} else {
			if len(currentNumberLocations) > 0 {
				val, _ := strconv.Atoi(currentNumber)
				coord := numberCoordinate{locations: currentNumberLocations, value: val}
				parsedNumbers = append(parsedNumbers, coord)

				currentNumber = ""
				currentNumberLocations = nil
			}
			index++
		}

	}

	// Adding this if statement took me half a day 🔥💩
	if len(currentNumberLocations) > 0 {
		val, _ := strconv.Atoi(currentNumber)
		coord := numberCoordinate{locations: currentNumberLocations, value: val}
		parsedNumbers = append(parsedNumbers, coord)
	}

	return parsedNumbers
}

func LineParser(line string, lineNumber int) []interface{} {
	var parsedLineItems []interface{}

	symbols := getSymbols(line, lineNumber)

	if len(symbols) > 0 {
		parsedLineItems = append(parsedLineItems, symbols...)
	}

	numbers := getNumbers(line, lineNumber)

	if len(numbers) > 0 {
		parsedLineItems = append(parsedLineItems, numbers...)
	}

	return parsedLineItems
}

func getTotalValueOfpartNumbers(numbers []numberCoordinate, symbolMap map[point]rune) int {
	total := 0

	for _, number := range numbers {
		var adjacentPoints []point

		for _, cell := range number.locations {

			adjacentPoints = append(adjacentPoints, cell.getAdjacentCells()...)
		}

		isPartNumber := false

		for _, adjacentPoint := range adjacentPoints {
			if _, ok := symbolMap[adjacentPoint]; ok {
				isPartNumber = true
				break
			}
		}

		if isPartNumber {
			total += number.value
		}
	}

	return total
}

func createSymbolMap(symbols []coordinate) map[point]rune {
	m := make(map[point]rune)

	for _, symbol := range symbols {
		m[symbol.location] = symbol.value
	}

	return m
}

func main() {
	parsedValues := utils.ParseInput(LineParser, "input.txt")

	var numbers []numberCoordinate
	var symbols []coordinate

	for _, value := range parsedValues {

		if convertedSymbol, ok := value.(coordinate); ok {
			symbols = append(symbols, convertedSymbol)
		}

		if convertedNumber, ok := value.(numberCoordinate); ok {
			numbers = append(numbers, convertedNumber)
		}
	}

	symbolMap := createSymbolMap(symbols)

	total := getTotalValueOfpartNumbers(numbers, symbolMap)

	fmt.Printf("%v\n", total)
}
