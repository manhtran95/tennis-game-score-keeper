package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	This program simulates a score keeping system for a tennis game.
*/

const (
	// list of points
	ZERO      = "0"
	FIFTHTEEN = "15"
	THIRTY    = "30"
	FORTY     = "40"
	ADVANCED  = "A"
	// list of states
	STATE_NORMAL = "StateNormal"
	STATE_DEUCE  = "StateDeuce"
	STATE_END    = "StateEnd"
)

var display_mappings = map[int]string{
	0: ZERO,
	1: FIFTHTEEN,
	2: THIRTY,
	3: FORTY,
	4: ADVANCED,
}

type Player struct {
	name string
	// score can be 0, 1, 2, 3 or 4, which is mapped to ZERO, FIFTHTEEN, THIRTY, FORTY and ADVANCED respectively
	score int
}

func main() {
	// loop for games
	for {
		/* initialize for the game
		gameState: state of the game, either STATE_NORMAL, STATE_DEUCE or STATE_END
		playerWinning, playerLosing: pointer to player winning or losing a particular point
		winner: pointer to player winning the game
		player1: first player of the game
		player2: second player of the game
		*/
		gameState := STATE_NORMAL
		var playerWinning, playerLosing, winner *Player
		player1 := Player{
			name:  "Player1",
			score: 0,
		}
		player2 := Player{
			name:  "Player2",
			score: 0,
		}

		fmt.Println("Tennis game started.")
		inputPlayerNames(&player1, &player2)

		// loop for points during the game
		for gameState != STATE_END {
			displayScore(player1, player2)
			playerWinning, playerLosing = processPointWinning(&player1, &player2)

			switch gameState {
			case STATE_NORMAL:
				gameState, winner = processStateNormal(playerWinning, playerLosing)
			case STATE_DEUCE:
				gameState, winner = processStateDeuce(playerWinning, playerLosing)
			}
		}

		playing := processGameEnding(winner)
		if !playing {
			break
		}
	}
	fmt.Println("Good bye!")
}

func processGameEnding(winner *Player) bool {
	var input string
	fmt.Printf("Congratz, *%s* has won the game!\n", winner.name)
	fmt.Println(`New game?(input 'y' or 'n')`)
	fmt.Scanf("%s", &input)
	if input != "y" {
		return false
	} else {
		fmt.Println("-----")
		return true
	}
}

func displayScore(player1 Player, player2 Player) {
	fmt.Printf("(score) %s - %s: %s - %s\n", player1.name, player2.name, display_mappings[player1.score], display_mappings[player2.score])
}

func inputPlayerNames(player1 *Player, player2 *Player) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please input Player 1 name:")
	scanner.Scan()
	player1.name = scanner.Text()
	fmt.Println("Please input Player 2 name:")
	scanner.Scan()
	player2.name = scanner.Text()
}

func processPointWinning(player1 *Player, player2 *Player) (*Player, *Player) {
	var input string
	for {
		fmt.Printf("Who wins this score?(input '1' for %s or '2' for %s)\n", player1.name, player2.name)
		fmt.Scanf("%s", &input)
		if input != "1" && input != "2" {
			fmt.Println(`Invalid input! Please try again!`)
			continue
		}
		if input == "1" {
			return player1, player2
		} else {
			return player2, player1
		}
	}
}

func processStateNormal(playerWinning *Player, playerLosing *Player) (string, *Player) {
	// win case
	if playerWinning.score == 3 {
		return STATE_END, playerWinning
	} else if playerWinning.score == 2 && playerLosing.score == 3 {
		// pre-deuce case
		playerWinning.score = 3
		return STATE_DEUCE, nil
	} else {
		// other case
		playerWinning.score++
		return STATE_NORMAL, nil
	}
}

func processStateDeuce(playerWinning *Player, playerLosing *Player) (string, *Player) {
	switch playerWinning.score {
	case 3:
		if playerLosing.score == 3 {
			playerWinning.score = 4
		} else {
			playerLosing.score = 3
		}
		return STATE_DEUCE, nil
	case 4:
		return STATE_END, playerWinning
	default:
		return "", nil
	}
}
