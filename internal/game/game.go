package game

import (
	"fmt"
	"log/slog"
)

type Game struct {
	PlayerA     *PlayerBoard
	PlayerB     *PlayerBoard
	Winner      string
	PlayerATurn bool
	IsOver      bool
}

func NewGame(pA, pB *PlayerBoard) *Game {
	return &Game{
		PlayerA:     pA,
		PlayerB:     pB,
		Winner:      "",
		PlayerATurn: true,
	}
}

func (g *Game) Turn(x, y uint8) (turnResult TurnResult, err error) {
	if g.PlayerATurn {
		turnResult, err = g.makeTurn(g.PlayerB, g.PlayerA.Name, x, y)
	} else {
		turnResult, err = g.makeTurn(g.PlayerA, g.PlayerB.Name, x, y)
	}
	return turnResult, err
}

func (g *Game) makeTurn(pBoard *PlayerBoard, attacker string, x, y uint8) (TurnResult, error) {
	slog.Info("Strike!", "Striker", attacker, "Enemy", pBoard.Name, "x", x, "y", y)
	g.PlayerATurn = !g.PlayerATurn
	turnResult := TurnResult{}
	hit, err := pBoard.Strike(x, y)
	if err != nil {
		return turnResult, err
	}
	turnResult.IsHit = hit
	if pBoard.IsDefeated() {
		turnResult.GameOver = true
		g.IsOver = true
		g.Winner = attacker
	}
	return turnResult, nil
}

type TurnResult struct {
	IsHit    bool
	GameOver bool
}

const totalHitsRequired = 5 + 4 + 3*2 + 2

type PlayerBoard struct {
	Name  string
	Hits  int
	Board [][]bool
}

type InitBoardParams struct {
	PlayerId string
	Ships    []Ship
	// Carrier    5 size
	// Battleship 4 size
	// Destroyer  3 size
	// Submarine  3 size
	// PatrolBoat 2 size
}

type Ship struct {
	Size       uint8
	Horizontal bool
	X          uint8
	Y          uint8
}

func (init *InitBoardParams) IsValid() bool {
	return true
}

func NewPlayerBoard(params *InitBoardParams) (*PlayerBoard, error) {
	if !params.IsValid() {
		return &PlayerBoard{}, fmt.Errorf("Invalid board")
	}
	board := make([][]bool, 10)
	for i := 0; i < 10; i++ {
		board[i] = make([]bool, 10)
	}
	// var id uint8 = 1
	for _, ship := range params.Ships {
		err := insertShip(&ship, &board)
		// id += 1
		if err != nil {
			return &PlayerBoard{Name: params.PlayerId, Board: board, Hits: 0}, err
		}
	}

	return &PlayerBoard{Name: params.PlayerId, Board: board, Hits: 0}, nil
}

func insertShip(ship *Ship, board *[][]bool) error {
	slog.Debug("Inserting ship", "ship", *ship)
	if ship.Horizontal {
		x := ship.X
		for i := 0; i < int(ship.Size); i++ {
			if (*board)[ship.Y][x] {
				return fmt.Errorf("Another ship on params x=%d, y=%x", x, ship.Y)
			}
			(*board)[ship.Y][x] = true
			x += 1
		}
	} else {
		y := ship.Y
		for i := 0; i < int(ship.Size); i++ {
			if (*board)[y][ship.X] {
				return fmt.Errorf("Another ship on params x=%d, y=%x", ship.X, y)
			}
			(*board)[y][ship.X] = true
			y += 1
		}
	}
	return nil
}

func (p *PlayerBoard) PrintBoard() {
	fmt.Printf("Board of %s\n", p.Name)
	for _, row := range p.Board {
		for _, val := range row {
			fmt.Printf("%6t ", val)
		}
		fmt.Println()
	}
}

func (p *PlayerBoard) Strike(x, y uint8) (bool, error) {
	if x < 0 || y < 0 || x > 9 || y > 9 {
		return false, fmt.Errorf("Out of bound strike.")
	}
	if !p.Board[x][y] {
		return false, nil
	}
	p.Board[x][y] = false
	p.Hits += 1
	return true, nil
}

func (p *PlayerBoard) IsDefeated() bool {
	return p.Hits == totalHitsRequired
}
