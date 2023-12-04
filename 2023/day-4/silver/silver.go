package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"utils"
)

type scratchCard struct {
	id             int
	winningNumbers []int
	cardNumbers    []int
}

func intArrayContains(array []int, item int) bool {
	for _, value := range array {
		if value == item {
			return true
		}
	}

	return false
}

func (card scratchCard) calculateScore() int {
	winningNumberCount := 0

	for _, number := range card.cardNumbers {
		if intArrayContains(card.winningNumbers, number) {
			winningNumberCount++
		}
	}

	score := 0

	if winningNumberCount > 0 {
		score++
		winningNumberCount--
	}

	for i := winningNumberCount; i > 0; i-- {
		score = score * 2
	}

	return score
}

func getCardId(input string) int {
	cardIdInfo := strings.TrimPrefix(strings.Split(input, ":")[0], "Card ")

	id, _ := strconv.Atoi(cardIdInfo)
	return id
}

func extractNumbersFromString(input string, delimiter string) []int {
	var numbers []int

	numberStrings := strings.Split(input, delimiter)

	for _, numberString := range numberStrings {
		number, _ := strconv.Atoi(numberString)

		numbers = append(numbers, number)
	}

	return numbers
}

func LineParser(line string, lineNumber int) []interface{} {
	var parsedLineItems []interface{}

	spaceRegex := regexp.MustCompile(`\s+`)

	cleanedString := spaceRegex.ReplaceAllLiteralString(line, " ")

	id := getCardId(cleanedString)

	winningNumbers := extractNumbersFromString(strings.Split(strings.Split(cleanedString, " | ")[0], ": ")[1], " ")

	numbers := extractNumbersFromString(strings.Split(cleanedString, " | ")[1], " ")

	parsedScratchCard := scratchCard{id: id, winningNumbers: winningNumbers, cardNumbers: numbers}

	parsedLineItems = append(parsedLineItems, parsedScratchCard)

	return parsedLineItems
}

func calculateTotalScore(scratchCards []scratchCard) int {
	score := 0

	for _, card := range scratchCards {
		score = score + card.calculateScore()
	}

	return score
}

func main() {
	input := utils.ParseInput(LineParser, "input.txt")

	var scratchCards []scratchCard

	for _, value := range input {
		if convertedScratchCard, ok := value.(scratchCard); ok {
			scratchCards = append(scratchCards, convertedScratchCard)
		}
	}

	totalScore := calculateTotalScore(scratchCards)

	fmt.Printf("%v\n", totalScore)
}
