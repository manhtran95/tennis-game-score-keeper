package main

import "fmt"

/*
	This program simulates a score keeping system for a tennis game.
*/

const 
(
	// list of points
	ZERO = "0"
	FIFTHTEEN = "15"
	THIRTY = "30"
	FORTY = "40"
	ADVANCED = "A"
	// list of states
	STATE_NORMAL = "StateNormal"
	STATE_DEUCE = "StateDeuce"
	STATE_END = "StateEnd"
)

var normal_map = map[string]string{
	ZERO: FIFTHTEEN,
	FIFTHTEEN: THIRTY,
	THIRTY: FORTY,
}

type Player struct {
	name string
	score string
}

func displayScore(player1 Player, player2 Player) {
	fmt.Printf("(score) %s - %s: %s - %s\n", player1.name, player2.name, player1.score, player2.score)
}

func main() {
	// loop for games
	for{
		// initialize for the game
		var input string
		gameState := STATE_NORMAL
		var playerWinning, playerLosing, winner *Player
		player1 := Player{
			name: "Player1",
			score: ZERO,
		}
		player2 := Player{
			name: "Player2",
			score: ZERO,
		}
		fmt.Println("Tennis game started.")
		fmt.Println("Please input Player 1 name:")
		fmt.Scanf("%s", &player1.name)
		fmt.Println("Please input Player 2 name:")
		fmt.Scanf("%s", &player2.name)
		// loop for points during a game
		for ;gameState != STATE_END ;{
			displayScore(player1, player2)
			// process input for point winner
			fmt.Printf("Who wins this score?(input '1' for %s or '2' for %s)\n", player1.name, player2.name)
			fmt.Scanf("%s", &input)
			if input != "1" && input != "2" {
				fmt.Println(`Invalid input! Please try again!`)
				continue
			}
			if input == "1" {
				playerWinning = &player1
				playerLosing = &player2
			} else {
				playerWinning, playerLosing = &player2, &player1
			}
			// check for gameState
			switch gameState {
			case STATE_NORMAL:
				// win case
				if playerWinning.score == FORTY {
					gameState = STATE_END
					winner = playerWinning
				} else if playerWinning.score == THIRTY && playerLosing.score == FORTY {
					// pre-deuce case
					gameState = STATE_DEUCE
					playerWinning.score = FORTY
				} else {
					// other case
					playerWinning.score = normal_map[playerWinning.score]
				}
			case STATE_DEUCE:
				switch playerWinning.score{
				case FORTY:
					if playerLosing.score == FORTY {
						playerWinning.score = ADVANCED
					} else {
						playerLosing.score = FORTY
					}
				case ADVANCED:
					gameState = STATE_END
					winner = playerWinning
				}
			}	
		}
		// process end of game
		fmt.Printf("Congratz, *%s* has won the game!\n", winner.name)
		fmt.Println(`New game?(input 'y' or 'n')`)
		fmt.Scanf("%s", &input)
		if input != "y" {
			fmt.Println("Good bye!")
			break
		}
		fmt.Println("-----")
	}
}