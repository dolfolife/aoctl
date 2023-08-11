package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	Rock    int = 1
	Paper       = 2
	Sissors     = 3
	Win         = 6
	Draw        = 3
	Lost        = 0
)

/*
 A = Rock from Player1
 B = Paper from Player1
 C = Sissors from Player1

 X = Rock from Player2
 Y = Paper from Player2
 Z = Sissors from Player2

 Calculate total score of the player2
*/
func calculateHandPlayPart1(play []string) (int, error) {
	player1 := play[0]
	player2 := play[1]

	switch player1 {
	case "A":
		switch player2 {
		case "X":
			return Draw + Rock, nil
		case "Y":
			return Win + Paper, nil
		case "Z":
			return Lost + Sissors, nil
		default:
			return 0, fmt.Errorf("error: Hand %s not found", player2)
		}
	case "B":
		switch player2 {
		case "X":
			return Lost + Rock, nil
		case "Y":
			return Draw + Paper, nil
		case "Z":
			return Win + Sissors, nil
		default:
			return 0, fmt.Errorf("error: Hand %s not found", player2)
		}
	case "C":
		switch player2 {
		case "X":
			return Win + Rock, nil
		case "Y":
			return Lost + Paper, nil
		case "Z":
			return Draw + Sissors, nil
		default:
			return 0, fmt.Errorf("error: Hand %s not found", player2)
		}
	default:
		return 0, fmt.Errorf("error: Hand %s not found", player1)
	}
}

/*
 A = Rock from Player1
 B = Paper from Player1
 C = Sissors from Player1

 X = Loose the Match
 Y = Draw the Match
 Z = Win the Match

 Calculate total score of the player2
*/
func calculateHandPlayPart2(play []string) (int, error) {
	player1 := play[0]
	desiredOutput := play[1]

	switch player1 {
	case "A":
		switch desiredOutput {
		case "X":
			return Lost + Sissors, nil
		case "Y":
			return Draw + Rock, nil
		case "Z":
			return Win + Paper, nil
		default:
			return 0, fmt.Errorf("error: Hand %s not found", desiredOutput)
		}
	case "B":
		switch desiredOutput {
		case "X":
			return Lost + Rock, nil
		case "Y":
			return Draw + Paper, nil
		case "Z":
			return Win + Sissors, nil
		default:
			return 0, fmt.Errorf("error: Hand %s not found", desiredOutput)
		}
	case "C":
		switch desiredOutput {
		case "X":
			return Lost + Paper, nil
		case "Y":
			return Draw + Sissors, nil
		case "Z":
			return Win + Rock, nil
		default:
			return 0, fmt.Errorf("error: Hand %s not found", desiredOutput)
		}
	default:
		return 0, fmt.Errorf("error: Hand %s not found", player1)
	}
}

func main() {

	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("error processing the file")
	}

	plays := strings.Split(string(b), "\n")

	totalPart1 := 0
	totalPart2 := 0
	for i := 0; i < len(plays)-1; i++ {
		playHands := strings.Split(plays[i], " ")
		score1, err1 := calculateHandPlayPart1(playHands)
		score2, err2 := calculateHandPlayPart2(playHands)

		if err1 != nil || err2 != nil {
			fmt.Printf("the guide of plays is corrupted in line %d: \n %s", i, err)
			os.Exit(1)
		}
		totalPart1 = totalPart1 + score1
		totalPart2 = totalPart2 + score2
	}

	fmt.Printf("Score of following Part 1 %d\n", totalPart1)
	fmt.Printf("Score of following Part 2 %d\n", totalPart2)
}
