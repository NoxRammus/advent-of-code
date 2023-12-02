package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

type game struct {
	red   int
	green int
	blue  int
}

func newGame(red int, green int, blue int) *game {
	g := game{red: red, green: green, blue: blue}

	return &g
}

type gameRecord struct {
	id    int
	games []game
}

func newGameRecord(id int, games []game) *gameRecord {
	gr := gameRecord{id: id, games: games}

	return &gr
}

func parseInput() []gameRecord {
	file, err := os.Open("./input.txt")
	checkError(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var games []gameRecord

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ":")

		gameIdentifier64, _ := strconv.ParseInt(strings.TrimLeft(splitLine[0], "Game "), 0, 32)
		gameIdentifier := int(gameIdentifier64)

		gameInfos := strings.Split(splitLine[1], ";")

		var gameHints []game

		for _, gameHint := range gameInfos {
			pieces := strings.Split(gameHint, ",")

			var red, blue, green int

			for _, hintColor := range pieces {
				hint := strings.Split(strings.Trim(hintColor, " "), " ")

				number64, _ := strconv.ParseInt(hint[0], 0, 32)
				number := int(number64)
				color := hint[1]

				if strings.Compare("red", color) == 0 {
					red = number
					continue
				}

				if strings.Compare("blue", color) == 0 {
					blue = number
					continue
				}

				if strings.Compare("green", color) == 0 {
					green = number
					continue
				}
			}

			game := newGame(red, green, blue)
			gameHints = append(gameHints, *game)
		}

		gameRecord := newGameRecord(gameIdentifier, gameHints)
		games = append(games, *gameRecord)
	}

	return games
}

func getMinimumSetOfCubes(gameRecords []gameRecord) []game {
	var minimumGames []game

	for _, recordedGame := range gameRecords {
		var red = 0
		var green = 0
		var blue = 0

		for _, individualGame := range recordedGame.games {
			if individualGame.red > red {
				red = individualGame.red
			}

			if individualGame.green > green {
				green = individualGame.green
			}

			if individualGame.blue > blue {
				blue = individualGame.blue
			}
		}

		minGame := newGame(red, green, blue)

		minimumGames = append(minimumGames, *minGame)
	}

	return minimumGames
}

func calculatePowerValue(games []game) int {
	total := 0
	for _, game := range games {
		gamePowerValue := game.red * game.green * game.blue

		total += gamePowerValue
	}

	return total
}

func main() {
	games := parseInput()

	minGameSet := getMinimumSetOfCubes(games)

	powerLevel := calculatePowerValue(minGameSet)

	fmt.Printf("%v", powerLevel)
}
