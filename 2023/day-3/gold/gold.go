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

type numberWrapper struct {
	id    int
	value int
}

type gear struct {
	location point
	ratio    int
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

func createNumberMap(numbers []numberCoordinate) map[point]numberWrapper {
	numberMap := make(map[point]numberWrapper)

	for id, numCoord := range numbers {
		wrapper := numberWrapper{id: id, value: numCoord.value}

		for _, location := range numCoord.locations {
			numberMap[location] = wrapper
		}
	}

	return numberMap
}

func (cell point) getAdjacentNumberCells(numberMap map[point]numberWrapper) []numberWrapper {
	var adjacentNumberWrappers []numberWrapper

	adjacentCells := cell.getAdjacentCells()

	for _, adjacentCell := range adjacentCells {
		wrapper, ok := numberMap[adjacentCell]

		if ok {
			wrapperId := wrapper.id

			shouldAddWrapper := true

			for _, adjacentWrapper := range adjacentNumberWrappers {
				if adjacentWrapper.id == wrapperId {
					shouldAddWrapper = false
					break
				}
			}

			if shouldAddWrapper {
				adjacentNumberWrappers = append(adjacentNumberWrappers, wrapper)
			}
		}
	}

	return adjacentNumberWrappers
}

func getAllGears(symbols []coordinate, numberMap map[point]numberWrapper) []gear {
	var gears []gear

	for _, symbol := range symbols {
		adjacentNumberCells := symbol.location.getAdjacentNumberCells(numberMap)

		if len(adjacentNumberCells) == 2 {
			foundGear := gear{location: symbol.location, ratio: adjacentNumberCells[0].value * adjacentNumberCells[1].value}

			gears = append(gears, foundGear)
		}
	}

	return gears
}

func sumGears(gears []gear) int {
	total := 0

	for _, gear := range gears {
		total += gear.ratio
	}

	return total
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

	numberMap := createNumberMap(numbers)

	gears := getAllGears(symbols, numberMap)

	total := sumGears(gears)

	fmt.Printf("%v\n", total)
}
