package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const maxSIZE int = 9

type theBoard struct {
	entry       [maxSIZE][maxSIZE]int
	playerScore int
	playerName  string
}

func printBoard(gameboard *theBoard) {

	for i := 0; i < maxSIZE; i++ {
		for j := 0; j < maxSIZE; j++ {
			fmt.Printf("%d ", gameboard.entry[i][j])
		}
		fmt.Printf("\n")
	}
}

func boardAddToken(gameboard *theBoard, X, Y int) error {

	/*
		if err != nil {
			fmt.Println("Enter an index within range\n")
		} else {
	*/
	gameboard.entry[X][Y] = 1
	gameboard.playerScore += 10
	//}
	return errors.New("IndexOutOfRange")
}

func main() {

	var gameboard *theBoard = new(theBoard)
	gameboard.playerScore = 0
	gameboard.playerName = "Guest"

	var choice, X, Y int

	for {
		fmt.Println("\n\nWelcome to GO\n")
		fmt.Println("Please make your selection\n")
		fmt.Println("1) Enter a Move\n")
		fmt.Println("2) Print the Board\n")
		fmt.Println("3) Check Player Score\n")
		fmt.Println("4) Enter Username\n")
		fmt.Println("5) Exit GO\n")
		fmt.Scan(&choice)
		fmt.Printf("\nchoice entered: <%d>\n\n", choice)

		switch choice {
		case 1:
			fmt.Println("Enter a coordinate for X: ")
			fmt.Scan(&X)
			fmt.Println("Enter a coordinate for Y: ")
			fmt.Scan(&Y)
			boardAddToken(gameboard, X, Y)
			fmt.Printf("\nmove entered: (%d, %d)\n", X, Y)
			printBoard(gameboard)
		case 2:
			printBoard(gameboard)
		case 3:
			fmt.Printf("Hello %s, your current score is: %d", gameboard.playerName, gameboard.playerScore)
		case 4:
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Please enter username: ")
			text, _ := reader.ReadString('\n')
			fmt.Scanln(&text)
			fmt.Printf("Your name is %s", text)
			gameboard.playerName = text
		case 5:
			fmt.Printf("%s, thanks for playing GO, your final score is %d :)", gameboard.playerName, gameboard.playerScore)
			os.Exit(0)
		}

	}

}
