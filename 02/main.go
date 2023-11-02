package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	RockPoints     = 1
	PaperPoints    = 2
	ScissorsPoints = 3
)

const (
	WinPoints  = 6
	DrawPoints = 3
	LosePoints = 0
)

type Points struct {
	Rock     int
	Paper    int
	Scissors int
}

type Result struct {
	Win  int
	Lose int
	Draw int
}

type Rules map[string]map[string]int

func NewResultRules() Rules {
	return Rules{
		"A": {
			"Y": DrawPoints + RockPoints,
			"Z": WinPoints + PaperPoints,
			"X": LosePoints + ScissorsPoints,
		},
		"B": {
			"X": LosePoints + RockPoints,
			"Y": DrawPoints + PaperPoints,
			"Z": WinPoints + ScissorsPoints,
		},
		"C": {
			"Z": WinPoints + RockPoints,
			"X": LosePoints + PaperPoints,
			"Y": DrawPoints + ScissorsPoints,
		},
	}
}

func resolveBattle(b string, r *int) {
	rules := NewResultRules()
	p1, p2 := splitMove(b)

	moveResult, ok := rules[p1][p2]
	if !ok {
		log.Fatalf("Invalid battle type: %s", b)
	}
	*r += moveResult
}

func splitMove(s string) (string, string) {
	battle := strings.Split(s, " ")
	if len(battle) != 2 {
		log.Fatalf("Invalid battle: %s", battle)
	}
	return battle[0], battle[1]
}

func main() {
	result := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		resolveBattle(scanner.Text(), &result)
	}
	fmt.Println(result)
}
