package game_test

import (
	"battleships/internal/game"
	"fmt"
	"math/rand/v2"
	"testing"
)

// CreateInitParamsMassive generates a testing board
// ┌───┬┬────┬────┬────┬────┬────┬────┬────┬────┬────┬────┐
//
// │ X ││ 1  │ 2  │ 3  │ 4  │ 5  │ 6  │ 7  │ 8  │ 9  │ 10 │
//
// ├───┼┼────┼────┼────┼────┼────┼────┼────┼────┼────┼────┤
//
// │ A ││ C1 │ C1 │ C1 │ C1 │ C1 │    │ S1 │    │ P4 │ D3 │
//
// │ B ││    │    │    │    │    │    │ S2 │    │ P4 │ D3 │
//
// │ C ││    │    │    │    │    │    │ S2 │    │    │ D3 │
//
// │ D ││ P3 │ B1 │ B1 │ B1 │ B1 │    │    │    │ P5 │ D1 │
//
// │ E ││ P3 │    │    │    │    │    │    │    │ P5 │ D1 │
//
// │ G ││    │    │    │    │    │    │    │    │    │ D1 │
//
// │ F ││    │    │    │    │ B2 │    │    │    │    │ D2 │
//
// │ H ││    │ P1 │    │    │ B2 │ S2 │ S2 │ S2 │    │ D2 │
//
// │ I ││    │ P1 │    │    │ B2 │ S3 │ S3 │ S3 │    │ D2 │
//
// │ J ││ P2 │ P2 │    │    │ B2 │    │ S4 │ S4 │ S4 │    │
//
// └───┴────┴────┴────┴────┴────┴────┴────┴────┴────┴────┘
func CreateInitParamsMassive() *game.InitBoardParams {
	ships := make([]game.Ship, 0, 15)
	ships = append(ships, game.Ship{Size: 5, Horizontal: true, X: 0, Y: 0})
	ships = append(ships, game.Ship{Size: 4, Horizontal: true, X: 1, Y: 3})
	ships = append(ships, game.Ship{Size: 4, Horizontal: false, X: 4, Y: 6})
	ships = append(ships, game.Ship{Size: 3, Horizontal: false, X: 9, Y: 3})
	ships = append(ships, game.Ship{Size: 3, Horizontal: false, X: 9, Y: 6})
	ships = append(ships, game.Ship{Size: 3, Horizontal: false, X: 9, Y: 0})
	ships = append(ships, game.Ship{Size: 3, Horizontal: false, X: 6, Y: 0})
	ships = append(ships, game.Ship{Size: 3, Horizontal: true, X: 5, Y: 7})
	ships = append(ships, game.Ship{Size: 3, Horizontal: true, X: 5, Y: 8})
	ships = append(ships, game.Ship{Size: 3, Horizontal: true, X: 6, Y: 9})
	ships = append(ships, game.Ship{Size: 2, Horizontal: false, X: 1, Y: 7})
	ships = append(ships, game.Ship{Size: 2, Horizontal: true, X: 0, Y: 9})
	ships = append(ships, game.Ship{Size: 2, Horizontal: false, X: 0, Y: 3})
	ships = append(ships, game.Ship{Size: 2, Horizontal: false, X: 8, Y: 0})
	ships = append(ships, game.Ship{Size: 2, Horizontal: false, X: 8, Y: 3})
	return &game.InitBoardParams{
		PlayerId: "TestUser",
		Ships:    ships,
	}
}

func TestBoardStrike(t *testing.T) {
	initParams := CreateInitParams("TestUser")
	fmt.Printf("%#v\n", *initParams)
	playerBoard, err := game.NewPlayerBoard(initParams)
	if err != nil {
		playerBoard.PrintBoard()
		t.Error(err)
	}
	playerBoard.PrintBoard()
	res, err := playerBoard.Strike(0, 0)
	if err != nil {
		playerBoard.PrintBoard()
		t.Error(err)
	}
	if !res {
		t.Fail()
	}
	fmt.Println()
	playerBoard.PrintBoard()
}

// ┌───┬────┬────┬────┬────┬────┬───┬────┬───┬───┬────┐
//
// │ X │ 1  │ 2  │ 3  │ 4  │ 5  │ 6 │ 7  │ 8 │ 9 │ 10 │
//
// ├───┼────┼────┼────┼────┼────┼───┼────┼───┼───┼────┤
//
// │ A │ C1 │ C1 │ C1 │ C1 │ C1 │   │ S1 │   │   │    │
//
// │ B │    │    │    │    │    │   │ S2 │   │   │    │
//
// │ C │    │    │    │    │    │   │ S2 │   │   │    │
//
// │ D │    │ B1 │ B1 │ B1 │ B1 │   │    │   │   │ D1 │
//
// │ E │    │    │    │    │    │   │    │   │   │ D1 │
//
// │ G │    │    │    │    │    │   │    │   │   │ D1 │
//
// │ F │    │    │    │    │    │   │    │   │   │    │
//
// │ H │    │ P1 │    │    │    │   │    │   │   │    │
//
// │ I │    │ P1 │    │    │    │   │    │   │   │    │
//
// │ J │    │    │    │    │    │   │    │   │   │    │
//
// └───┴────┴────┴────┴────┴────┴───┴────┴───┴───┴────┘
func CreateInitParams(id string) *game.InitBoardParams {
	ships := make([]game.Ship, 0, 15)
	ships = append(ships, game.Ship{Size: 5, Horizontal: true, X: 0, Y: 0})
	ships = append(ships, game.Ship{Size: 4, Horizontal: true, X: 1, Y: 3})
	ships = append(ships, game.Ship{Size: 3, Horizontal: false, X: 9, Y: 3})
	ships = append(ships, game.Ship{Size: 3, Horizontal: false, X: 6, Y: 0})
	ships = append(ships, game.Ship{Size: 2, Horizontal: false, X: 1, Y: 7})
	return &game.InitBoardParams{
		PlayerId: id,
		Ships:    ships,
	}
}

func TestSimulate(t *testing.T) {
	initA := CreateInitParams("playerA")
	initB := CreateInitParams("playerB")
	pA, _ := game.NewPlayerBoard(initA)
	pB, _ := game.NewPlayerBoard(initB)
	currentGame := game.NewGame(pA, pB)
	currentGame.PlayerA.PrintBoard()
	fmt.Println()
	currentGame.PlayerB.PrintBoard()
	for !currentGame.IsOver {
		currentGame.Turn(randomCord(), randomCord())
	}
	fmt.Printf("Player %s won\n", currentGame.Winner)
	currentGame.PlayerA.PrintBoard()
	fmt.Println()
	currentGame.PlayerB.PrintBoard()
}

func randomCord() uint8 {
	return uint8(rand.IntN(10))
}
