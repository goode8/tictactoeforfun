//Dawn Goodnight
//CSC500
//FA20
//Tic-Tac-Toe
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// allow [][]byte to implement the sort.Interface interface
type lexicographically [][]byte


// bytes.Compare compares the byte slices lexicographically (reverse alphabetically to put highest score at the top)
func (l lexicographically) Less(j, i int) bool { return bytes.Compare(l[i], l[j]) < 0 }
func (l lexicographically) Len() int           { return len(l) }
func (l lexicographically) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

func GetName() string {
	fmt.Println("Welcome to Tic-Tac-Toe")
	fmt.Println("My name is Gunilla")
	var fname string
	fmt.Println("What is your name?")
	fmt.Scanf("%s\n", &fname)
	fmt.Println("                         ")
	fmt.Println("Hello", fname)
	fmt.Println("                         ")
	return fname
}

func LevelofPlay() int {
	var level string
	fmt.Println("                         ")
	fmt.Println("                         ")
	fmt.Println("What level would you like to play?")
	fmt.Println("The options are: Auto, Random, Easy, Medium, or Difficult")
	fmt.Println("press 1 for computer vs. computer, 2 for random, 3 for easy, and 4 for medium, 5 for difficult")
	fmt.Println("then press enter or return")
	fmt.Println("                         ")
	fmt.Println("                         ")
	//fmt.Println("manual override to Easy for testing...")
	fmt.Scanf("%s\n", &level) //gets a string!
	intLevel, err10 := strconv.Atoi(level)
	err10 = nil
	if err10 != nil {
		panic(err10)
		//fmt.Println(intLevel)
	}
	for i := 0; i < 4; i++ { //five tries to get choice of level of play
		if intLevel < 6 && intLevel > 0 {
			return intLevel
			i = 10
		} else {
			fmt.Println("Please type a number from 1 to 4 and press enter or return")
			fmt.Scanf("%s\n", &level)
			intLevel, err10 := strconv.Atoi(level)
			err10 = nil
			if err10 != nil {
				panic(err10)
				//fmt.Println(intLevel)
			}
			if intLevel < 6 && intLevel > 0 {
				return intLevel
				break
			}
			//intLevel = 2
		}
	}

	return intLevel
}

func BoardInitialize() [9]string { //start to play ttt
	//Board1 := list.New()
	//var Board1 [9]string
	fmt.Println("                         ")
	fmt.Println("Let's Play!")
	fmt.Println("                         ")
	fmt.Println("                         ")
	Board1 := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	return Board1
}

func PrintBoard(Board1 [9]string) {
	fmt.Println("                         ")
	fmt.Println(Board1[0], "|", Board1[1], "|", Board1[2])
	fmt.Println("---------")
	fmt.Println(Board1[3], "|", Board1[4], "|", Board1[5])
	fmt.Println("---------")
	fmt.Println(Board1[6], "|", Board1[7], "|", Board1[8])
	fmt.Println("                         ")
}

func Getchoice1(Board1 [9]string) (string, [9]string, string) {
	fmt.Println("                         ")
	fmt.Println("                         ")
	fmt.Println("It's your turn.")
	fmt.Println("To choose where your X goes, type a number and press enter or return.")
	fmt.Println("                         ")
	var choice1 string
	fmt.Scanf("%s\n", &choice1)
	Board2 := Board1
	//borrowed code!
	for i := 0; i < 5; i++ { //five tries to get choice
		//value not location
		ilocation, found := Find(Board1, choice1)
		if !found {
			fmt.Println("Please type an available number from the grid and press enter or return")
			fmt.Scanf("%s\n", &choice1)
			var err5 error
			err5 = nil
			if err5 != nil {
				panic(err5)
				fmt.Print(ilocation)
			}
		}
	}
	thespot := "X"
	//	fmt.Println("Good Choice!", choice1, Board2)

	return choice1, Board2, thespot //value not location
}


func Find(board [9]string, choice string) (int, bool) {
	for ilocation, stringvalue := range board {
		if stringvalue == choice {
			return ilocation, true //return location not value
			//	fmt.Println("location:", ilocation)
			//	fmt.Println("found")
		}
	}
	//fmt.Println("not found")
	return -1, false

}

func PlayAgain() bool {
	fmt.Println("                         ")
	fmt.Println("                         ")
	fmt.Println("Would you like to play again?")
	fmt.Println("Press y and then enter or return to play again")
	fmt.Println("                         ")
	var choicetp string
	fmt.Scanf("%s\n", &choicetp)
	if choicetp == "y" || choicetp == "Y" {
		return true
	} else {
		//take that as a no
		return false
	}
}

func UpdateBoard(choicevalue string, Board [9]string, thespot string) [9]string {
	//value input (not location)
	//fmt.Println("updateboard")
	//fmt.Println("choicevalue", choicevalue)
	//fmt.Println("Board", Board)

	ichoicevalueint, err4 := strconv.Atoi(choicevalue)
	ichoicelocationint := ichoicevalueint - 1 //to location
	//fmt.Println("ichoicevalueint", ichoicevalueint)
	//fmt.Println("ichoicelocationint", ichoicelocationint)

	if err4 != nil {
		panic(err4)
	}
	Board[ichoicelocationint] = thespot
	//fmt.Println("board after", Board)
	return Board
}

func GunillaGoesRandom(Board [9]string) [9]string {
	fmt.Println("                         ")
	fmt.Println("Random Play.")
	fmt.Print("It's my turn.")
	fmt.Println("                         ")
	fmt.Print("Let me think...")
	//randomly pick a number from 1-9.//value not location
	var BoardNext [9]string
	BoardNext = Board
	//random number generator
	rand.Seed(time.Now().UnixNano())
	var RandomHelper3 int
	var RandomHelper4 string
	//check if it's available
	iteratori := 0
	for iteratori == 0 {
		rand.NewSource(time.Now().UnixNano())
		RandomHelper3 = randomInt(1, 10) //value not location
		RandomHelper4 = strconv.Itoa(RandomHelper3)
		//returns number and bool
		ilocation, found := Find(Board, RandomHelper4)
		var err7 error
		err7 = nil
		if err7 != nil {
			fmt.Print(ilocation)
			panic(err7)
		}
		if found { //true found, all good continue
			//fmt.Println("9 found", iteratori)
			//fmt.Errorf("success")
			iteratori = iteratori + 1
		} else {
			fmt.Println("Hmmmm, please give me a moment.")
		} //if

	} //for
	thespot := "O"
	BoardNext = UpdateBoard(RandomHelper4, Board, thespot) //value not location
	//	fmt.Println("RH4 ", RandomHelper4, "BN ", BoardNext)
	return BoardNext
}

func GunillaGoesEasy(Board [9]string) [9]string { //never block on purpose, use sides not corners
	//first turn, it's zero but subsequent visits it is higher.
	OhCount := ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O")))
	thespot := "O"
	fmt.Println("                         ")
	fmt.Println("Easy Play.")
	fmt.Print("It's my turn.")
	fmt.Println("                         ")
	fmt.Print("Let me think...")

	//all offensive moves
	if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
		i, j, k := 0, 1, 2
		Board = Offense(i, j, k, Board, thespot)
	}
	if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
		i, j, k := 3, 4, 5
		Board = Offense(i, j, k, Board, thespot)
	}
	if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
		i, j, k := 6, 7, 8
		Board = Offense(i, j, k, Board, thespot)
	}
	if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
		i, j, k := 0, 3, 6
		Board = Offense(i, j, k, Board, thespot)
	}
	if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
		i, j, k := 1, 4, 7
		Board = Offense(i, j, k, Board, thespot)
	}
	if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
		i, j, k := 2, 5, 8
		Board = Offense(i, j, k, Board, thespot)
	}
	if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
		i, j, k := 0, 4, 8
		Board = Offense(i, j, k, Board, thespot)
	}
	if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
		i, j, k := 2, 4, 6
		Board = Offense(i, j, k, Board, thespot)
	}

	//if no offensive moves

	EasyChoices := [9]string{"4", "6", "5", "8", "2", "9", "1", "3", "7"} //values not location
	//Easybylocation := [9]string{"1", "3", "5", "7", "0", "2", "6", "8", "4"} //values not location
	//check if it's available
	iteratori := 0 //giant for loop to keep playing
	iteratorA := 0 //location not value

	if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
		for iteratori == 0 {
			//returns number and bool
			ilocation, found := Find(Board, EasyChoices[iteratorA]) //ilocation=board location in Board, found=true if valid choice only. or -1, false
			//for ilocation, stringvalue := range board {
			var err7 error
			err7 = nil
			if err7 != nil {
				fmt.Print(ilocation)
				panic(err7)
			}
			if found { //true found, all good continue
				thespot := "O"

				var choice string
				choice = Board[ilocation] //4,6,5
				Board = UpdateBoard(choice, Board, thespot)
				iteratori = iteratori + 1 //get outta loop
			} else {
				//fmt.Println("                         ")
				fmt.Println("Hmmmm, gimme a moment")
				iteratorA = iteratorA + 1 //try next value in order
			} //if
		}
	}

	return Board
}

func GunillaGoesMedium(Board [9]string) [9]string { //block
	thespot := "O"
	fmt.Println("                         ")
	fmt.Println("Medium Play.")
	fmt.Print("It's my turn.")
	fmt.Println("                         ")
	fmt.Print("Let me think...")
	fmt.Println("                         ")
	fmt.Println("                         ")
	//first turn, it's zero but subsequent visits it is higher.
	OhCount := ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O")))
	//MedChoices := [9]string{"5", "9", "1", "8", "7", "3", "2", "4", "6"} //values not location

	//first turn only
	XhCount := (strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))
	if 1 == XhCount {
		if Board[4] == "5" {
			Board[4] = "O"
		} else if Board[8] == "9" { //middle not available
			//	fmt.Println("medium best choice ", MedChoices)
			Board[8] = "O"
		} //all other turns
	}
	XhCount = (strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))
	if 1 < XhCount {

		//all defensive moves
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 6, 7, 8
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 0, 3, 6
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 0, 1, 2
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 1, 4, 7
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 2, 5, 8
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 3, 4, 5
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 0, 4, 8
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 2, 4, 6
			Board = Defense(i, j, k, Board, thespot)
		}
	}

	if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {

		//play anything
		if Board[8] == "9" {
			Board[8] = "O"
		} else if Board[0] == "1" {
			Board[0] = "O"
		} else if Board[7] == "8" {
			Board[7] = "O"
			//"7", "3", "2", "4", "6"
		} else if Board[6] == "7" {
			Board[6] = "O"
		} else if Board[2] == "3" {
			Board[2] = "O"
		} else if Board[1] == "2" {
			Board[1] = "O"
		} else if Board[3] == "4" {
			Board[3] = "O"
		} else if Board[5] == "6" {
			Board[5] = "O"
		}
	}

	return Board
}

func GunillaGoesDifficult(Board [9]string) [9]string {
	thespot := "O"
	fmt.Println("                         ")
	fmt.Println("Difficult Play.")
	fmt.Print("It's my turn.")
	fmt.Println("                         ")
	fmt.Print("Let me think...")
	fmt.Println("                         ")
	fmt.Println("                         ")
	//first turn, it's zero but subsequent visits it is higher.
	OhCount := ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O")))
	//MedChoices := [9]string{"5", "9", "1", "8", "7", "3", "2", "4", "6"} //values not location

	//first turn only
	XhCount := (strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))
	if 1 == XhCount {
		if Board[4] == "5" {
			Board[4] = "O"
		} else if Board[8] == "9" { //middle not available
			//	fmt.Println("medium best choice ", MedChoices)
			Board[8] = "O"
		} //all other turns
	}
	XhCount = (strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))
	if 1 < XhCount {
		//all offensive moves
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 0, 1, 2
			Board = Offense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 3, 4, 5
			Board = Offense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 6, 7, 8
			Board = Offense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 0, 3, 6
			Board = Offense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 1, 4, 7
			Board = Offense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 2, 5, 8
			Board = Offense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 0, 4, 8
			Board = Offense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 2, 4, 6
			Board = Offense(i, j, k, Board, thespot)
		}
		//all defensive moves
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 6, 7, 8
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 0, 3, 6
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 0, 1, 2
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 1, 4, 7
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 2, 5, 8
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 3, 4, 5
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 0, 4, 8
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 2, 4, 6
			Board = Defense(i, j, k, Board, thespot)
		}
	}

	if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {

		//play anything
		if Board[8] == "9" {
			Board[8] = "O"
		} else if Board[0] == "1" {
			Board[0] = "O"
		} else if Board[6] == "7" {
			Board[6] = "O"
			//"7", "3", "2", "4", "6"
			//some optimization needed ?
		} else if Board[7] == "8" {
			Board[7] = "O"
		} else if Board[2] == "3" {
			Board[2] = "O"
		} else if Board[1] == "2" {
			Board[1] = "O"
		} else if Board[3] == "4" {
			Board[3] = "O"
		} else if Board[5] == "6" {
			Board[5] = "O"
		}
	}

	return Board
}

func GunillaVSGunilla1(Board [9]string) [9]string {
	thespot := "X"
	fmt.Println("                         ")
	fmt.Println("We love to play when we are alone together.")
	fmt.Print("It's my turn. (X)")
	fmt.Println("                         ")
	fmt.Print("This is my kind of fun.")
	fmt.Println("                         ")
	fmt.Println("                         ")
	//first turn, it's zero but subsequent visits it is higher.
	XXCount := ((strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X")))
	//first turn only
	if 0 == XXCount && Board[4] == "5" {
		Board[4] = "X"
	}
	//all other turns
	//	OOCount = (strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))
	//	if 1 < OOCount {

	//all offensive moves
	if XXCount == ((strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))) {
		i, j, k := 0, 1, 2
		Board = Offense(i, j, k, Board, thespot)
	}
	if XXCount == ((strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))) {
		i, j, k := 3, 4, 5
		Board = Offense(i, j, k, Board, thespot)
	}
	if XXCount == ((strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))) {
		i, j, k := 6, 7, 8
		Board = Offense(i, j, k, Board, thespot)
	}
	if XXCount == ((strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))) {
		i, j, k := 0, 3, 6
		Board = Offense(i, j, k, Board, thespot)
	}
	if XXCount == ((strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))) {
		i, j, k := 1, 4, 7
		Board = Offense(i, j, k, Board, thespot)
	}
	if XXCount == ((strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))) {
		i, j, k := 2, 5, 8
		Board = Offense(i, j, k, Board, thespot)
	}
	if XXCount == ((strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))) {
		i, j, k := 0, 4, 8
		Board = Offense(i, j, k, Board, thespot)
	}
	if XXCount == ((strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))) {
		i, j, k := 2, 4, 6
		Board = Offense(i, j, k, Board, thespot)
	}
	//all defensive moves
	if XXCount == ((strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))) {
		i, j, k := 6, 7, 8
		Board = Defense(i, j, k, Board, thespot)
	}
	if XXCount == ((strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))) {
		i, j, k := 0, 3, 6
		Board = Defense(i, j, k, Board, thespot)
	}
	if XXCount == ((strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))) {
		i, j, k := 0, 1, 2
		Board = Defense(i, j, k, Board, thespot)
	}
	if XXCount == ((strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))) {
		i, j, k := 1, 4, 7
		Board = Defense(i, j, k, Board, thespot)
	}
	if XXCount == ((strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))) {
		i, j, k := 2, 5, 8
		Board = Defense(i, j, k, Board, thespot)
	}
	if XXCount == ((strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))) {
		i, j, k := 3, 4, 5
		Board = Defense(i, j, k, Board, thespot)
	}
	if XXCount == ((strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))) {
		i, j, k := 0, 4, 8
		Board = Defense(i, j, k, Board, thespot)
	}
	if XXCount == ((strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))) {
		i, j, k := 2, 4, 6
		Board = Defense(i, j, k, Board, thespot)
	}

	if XXCount == ((strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))) {
		//no actions have happened play anything
		//play anything
		if Board[2] == "3" {
			Board[2] = "X"
		} else if Board[6] == "7" {
			Board[6] = "X"
		} else if Board[5] == "6" {
			Board[5] = "X"
		} else if Board[2] == "3" {
			Board[2] = "X"
		} else if Board[0] == "1" {
			Board[0] = "X"
		} else if Board[1] == "2" {
			Board[1] = "X"
		} else if Board[3] == "4" {
			Board[3] = "X"
		} else if Board[8] == "9" {
			Board[8] = "X"
		}
	}

	return Board
}

func GunillaVSGunilla2(Board [9]string) [9]string {
	thespot := "O"
	fmt.Println("                         ")
	fmt.Println("Love it!")
	fmt.Print("It's my turn (O).")
	fmt.Println("                         ")
	fmt.Print("Let me think...")
	fmt.Println("                         ")
	fmt.Println("                         ")
	//first turn, it's zero but subsequent visits it is higher.
	OhCount := ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O")))
	//MedChoices := [9]string{"5", "9", "1", "8", "7", "3", "2", "4", "6"} //values not location

	//first turn only
	XhCount := (strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))
	if 1 == XhCount {
		if Board[4] == "5" {
			Board[4] = "O"
		}
	}
	//all other turns
	XhCount = (strings.Count(Board[0], "X")) + (strings.Count(Board[1], "X")) + (strings.Count(Board[2], "X")) + (strings.Count(Board[3], "X")) + (strings.Count(Board[4], "X")) + (strings.Count(Board[5], "X")) + (strings.Count(Board[6], "X")) + (strings.Count(Board[7], "X")) + (strings.Count(Board[8], "X"))
	if 1 < XhCount {
		//all offensive moves
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 0, 1, 2
			Board = Offense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 3, 4, 5
			Board = Offense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 6, 7, 8
			Board = Offense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 0, 3, 6
			Board = Offense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 1, 4, 7
			Board = Offense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 2, 5, 8
			Board = Offense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 0, 4, 8
			Board = Offense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 2, 4, 6
			Board = Offense(i, j, k, Board, thespot)
		}
		//all defensive moves
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 6, 7, 8
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 0, 3, 6
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 0, 1, 2
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 1, 4, 7
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 2, 5, 8
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 3, 4, 5
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 0, 4, 8
			Board = Defense(i, j, k, Board, thespot)
		}
		if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {
			i, j, k := 2, 4, 6
			Board = Defense(i, j, k, Board, thespot)
		}
	}

	if OhCount == ((strings.Count(Board[0], "O")) + (strings.Count(Board[1], "O")) + (strings.Count(Board[2], "O")) + (strings.Count(Board[3], "O")) + (strings.Count(Board[4], "O")) + (strings.Count(Board[5], "O")) + (strings.Count(Board[6], "O")) + (strings.Count(Board[7], "O")) + (strings.Count(Board[8], "O"))) {

		//play anything
		if Board[8] == "9" {
			Board[8] = "O"
		} else if Board[0] == "1" {
			Board[0] = "O"
		} else if Board[6] == "7" {
			Board[6] = "O"
			//"7", "3", "2", "4", "6"
			//some optimization needed ?
		} else if Board[7] == "8" {
			Board[7] = "O"
		} else if Board[2] == "3" {
			Board[2] = "O"
		} else if Board[1] == "2" {
			Board[1] = "O"
		} else if Board[3] == "4" {
			Board[3] = "O"
		} else if Board[5] == "6" {
			Board[5] = "O"
		}
	}

	return Board
}

func Defense(i int, j int, k int, Board [9]string, thespot string) [9]string {
	//fmt.Println(i, j, k, Board)
	var otherSpot string
	if thespot == "O" {
		otherSpot = "X"
	} else {
		otherSpot = "O"
	}

	if (strings.Count(Board[i], otherSpot))+(strings.Count(Board[j], otherSpot))+(strings.Count(Board[k], otherSpot)) == 2 && ((strings.Count(Board[i], thespot))+(strings.Count(Board[j], thespot))+(strings.Count(Board[k], thespot)) == 0) {
		//	fmt.Println("insidecompareblock if loop")
		if Board[i] != otherSpot {
			Board[i] = thespot
		} else if Board[j] != otherSpot {
			Board[j] = thespot
		} else if Board[k] != otherSpot {
			Board[k] = thespot
		}
	}
	return Board
}

func Offense(i int, j int, k int, Board [9]string, thespot string) [9]string {
	//fmt.Println(i, j, k, Board)
	//playing some offense
	var otherSpot string
	if thespot == "O" {
		otherSpot = "X"
	} else {
		otherSpot = "O"
	}
	if ((strings.Count(Board[i], thespot))+(strings.Count(Board[j], thespot))+(strings.Count(Board[k], thespot)) == 2) && ((strings.Count(Board[i], otherSpot))+(strings.Count(Board[j], otherSpot))+(strings.Count(Board[k], otherSpot)) == 0) {
		//fmt.Println("insidecompareblock if loop")
		if Board[i] != thespot {
			Board[i] = thespot
		} else if Board[j] != thespot {
			Board[j] = thespot
		} else if Board[k] != thespot {
			Board[k] = thespot
		}
	}
	return Board
}

func NewHighScore(fname, newhighscore string) {
	//this is really writing...
	currentTime := time.Now()
	newlinescore := []byte(fmt.Sprintf("%v, %v, %v\n", newhighscore, fname, currentTime.Format("01-02-2006")))
	HSfile, err1 := os.OpenFile("highscore.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err1 != nil {
		panic(err1)
	}

	defer HSfile.Close()
	if _, err1 := HSfile.WriteString(string(newlinescore)); err1 != nil {
		log.Println(err1)
	}
}

func ScoreBoardPrint(Fname, newscore string) {
	//read, sort, and print

	existingscorelines, err2 := ioutil.ReadFile("highscore.txt")
	if err2 != nil {
		panic(err2)
	}
	//read per line
	singlelines := bytes.Split(existingscorelines, []byte{'\n'})
	//sort per line
	sort.Sort(lexicographically(singlelines))
	existingscorelines = bytes.Join(singlelines, []byte{'\n'})

	//print score board!!
	intscore, err3 := strconv.Atoi(newscore)
	if err3 != nil {
		panic(err3)
	}

	fmt.Println("                         ")
	fmt.Println("                         ")
	fmt.Println("********************Score Board!!********************")
	fmt.Println("                         ")
	if intscore > 0 {
		fmt.Println(strings.ToUpper(Fname), "your score is:  ", newscore, "   You Made The List!!")
		fmt.Println("                         ")
	} else {
		fmt.Println(Fname, "you didn't score. Better luck next time.")
		fmt.Println("                         ")
	}

	fmt.Println("SCORE****NAME********DATE*************")
	fmt.Print(string(existingscorelines))
	fmt.Println("                         ")
	fmt.Println("*************************************")
	//}
}

//borrowed code:
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func WonOrNot(board [9]string) bool {
	//var testvalue  [9]string
	//testvalue[0:8] = ["X", "X", "X", "", "", "", "", "", ""]
	if board[0] == board[1] && board[1] == board[2] || board[3] == board[4] && board[4] == board[5] || board[6] == board[7] && board[7] == board[8] || board[0] == board[3] && board[3] == board[6] || board[1] == board[4] && board[4] == board[7] || board[2] == board[5] && board[5] == board[8] || board[0] == board[4] && board[4] == board[8] || board[2] == board[4] && board[4] == board[6] {
		return true
	} else {
		return false
	}
	//value of board[0] wins!
}

func main() {
	Fname := GetName()

	playagain := true
	var newscore int
	var gameover int
	var Board3 [9]string
	var Board4 [9]string
	var Board5 [9]string
	var Board6 [9]string
	//var newhighscore string

	for playagain == true {
		Level := LevelofPlay()
		//	fmt.Println("Level: ", Level)
		Board1 := BoardInitialize()
		PrintBoard(Board1)
		if gameover = 0; gameover < 1 { //start or continue play!
			for iteratorj := 0; iteratorj < 4; iteratorj++ { //max no of turns - 1 if no one wins
				if Level == 1 {
					Board3 = GunillaVSGunilla1(Board1)
				} else {
					ChoiceA, Board2, thespot := Getchoice1(Board1) //value not location
					Board3 = UpdateBoard(ChoiceA, Board2, thespot)
				}
				PrintBoard(Board3)
				wonornot1 := WonOrNot(Board3)

				if wonornot1 == true {
					gameover++
					if Level == 1 {
						fmt.Println("I won against myself")
						fmt.Println("You were such a good sport, you get 300 points")
						newscore = newscore + 300
						break
					} else {
						fmt.Println("You won!")
						newscore = newscore + 100
						break
					}

				} else { //not won computer plays
					if Level == 1 {
						Board4 = GunillaVSGunilla2(Board3)
						Board1 = Board4 //to send back to top!
					} else if Level == 2 {
						Board4 = GunillaGoesRandom(Board3)
						Board1 = Board4 //to send back to top!
					} else if Level == 3 {
						Board4 = GunillaGoesEasy(Board3)
						Board1 = Board4 //to send back to top!
					} else if Level == 4 {
						Board4 = GunillaGoesMedium(Board3)
						Board1 = Board4 //to send back to top!
					} else { //if they don't choose a level default here!
						Board4 = GunillaGoesDifficult(Board3)
						Board1 = Board4 //to send back to top!
					}
				}
				PrintBoard(Board4)
				wonornot2 := WonOrNot(Board4)

				if wonornot2 == true {
					if Level == 1 {
						fmt.Println("I won against myself")
						fmt.Println("You were such a good sport, you get 300 points")
						newscore = newscore + 300
						break
					} else {
						gameover = gameover + 1
						fmt.Println("I won!")
						break
					}
				}
			}

			if gameover == 0 {
				//last turn
				//for filling all boxes if needed
				if Level == 1 {
					Board6 = GunillaVSGunilla1(Board1)
				} else {
					ChoiceA, Board5, thespot := Getchoice1(Board1)
					Board6 = UpdateBoard(ChoiceA, Board5, thespot)
				}
				PrintBoard(Board6)
				gameover = gameover + 1
				wonornot3 := WonOrNot(Board6)
				if wonornot3 == true {
					if Level == 1 {
						fmt.Println("I won against myself")
						fmt.Println("You were such a good sport, you get 300 points")
						newscore = newscore + 300
					} else {
						gameover = gameover + 1
						fmt.Println("You won!")
						newscore = newscore + 100
					}
				} else {
					CatsGame := [9]string{"C", "C", "C", "C", "C", "C", "C", "C", "C"}
					PrintBoard(CatsGame)
					fmt.Println("It's a draw. No points awarded. Good effort!")
				}
			} //end of last turn
		} //end of game over loop
		playagain = PlayAgain()
	}

	var errb5 error
	errb5 = nil
	if errb5 != nil {
		fmt.Println(Board5)
		panic(errb5)
	}

	//high scores recorded
	//padnewscore := fmt.Sprintf("%05d", newscore)
	newstringscore := strconv.Itoa(newscore)
	if newscore > 0 { //if scored anything write it down to the file
		NewHighScore(Fname, newstringscore)
	}
	ScoreBoardPrint(Fname, newstringscore)
}

/*
Project References

Get current date and time in various format in golang - golangprograms.com
https://www.golangprograms.com/get-current-date-and-time-in-various-format-in-golang.html

//random number generator
Go by Example: Random Numbers
https://gobyexample.com/random-numbers

//random number code:
Generating random numbers and strings in Go (flaviocopes.com)
https://flaviocopes.com/go-random/

*/
