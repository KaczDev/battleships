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
}

func NewGame(pA, pB *PlayerBoard) *Game {
	return &Game{
		PlayerA:     pA,
		PlayerB:     pB,
		Winner:      "",
		PlayerATurn: true,
	}
}

func (g *Game) Turn(x, y uint8, isPlayerA bool) (TurnResult, error) {
	if isPlayerA {
		return g.makeTurn(g.PlayerB, x, y)
	}
	return g.makeTurn(g.PlayerA, x, y)
}

func (g *Game) makeTurn(pBoard *PlayerBoard, x, y uint8) (TurnResult, error) {
	g.PlayerATurn = !g.PlayerATurn
	turnResult := TurnResult{}
	hit, err := pBoard.Strike(x, y)
	if err != nil {
		return turnResult, err
	}
	turnResult.IsHit = hit
	if pBoard.IsDefeated() {
		g.Winner = pBoard.Name
		turnResult.GameOver = true
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
	Board [][]uint8
}

type InitBoardParams struct {
	PlayerId string
	Ships    []Ship
	// Carrier    5 size, 1pc
	// Battleship 4 size, 2pc
	// Destroyer  3 size, 3pc
	// Submarine  3 size, 4pc
	// PatrolBoat 2 size, 5pc
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
	board := make([][]uint8, 10)
	for i := 0; i < 10; i++ {
		board[i] = make([]uint8, 10)
	}
	var id uint8 = 1
	for _, ship := range params.Ships {
		err := insertShip(&ship, &board, id)
		id += 1
		if err != nil {
			return &PlayerBoard{Name: params.PlayerId, Board: board, Hits: 0}, err
		}
	}

	return &PlayerBoard{Name: params.PlayerId, Board: board, Hits: 0}, nil
}

func insertShip(ship *Ship, board *[][]uint8, id uint8) error {
	slog.Debug("Inserting ship", "id", id, "ship", *ship)
	if ship.Horizontal {
		x := ship.X
		for i := 0; i < int(ship.Size); i++ {
			if (*board)[ship.Y][x] != 0 {
				return fmt.Errorf("Another ship on params x=%d, y=%x", x, ship.Y)
			}
			(*board)[ship.Y][x] = id
			x += 1
		}
	} else {
		y := ship.Y
		for i := 0; i < int(ship.Size); i++ {
			if (*board)[y][ship.X] != 0 {
				return fmt.Errorf("Another ship on params x=%d, y=%x", ship.X, y)
			}
			(*board)[y][ship.X] = id
			y += 1
		}
	}
	return nil
}

func (p *PlayerBoard) PrintBoard() {
	for _, row := range p.Board {
		for _, val := range row {
			fmt.Printf("%4d ", val)
		}
		fmt.Println()
	}
}

func (p *PlayerBoard) Strike(x, y uint8) (bool, error) {
	if x < 0 || y < 0 || x > 9 || y > 9 {
		return false, fmt.Errorf("Out of bound strike.")
	}
	slog.Info("Strike!", "x", x, "y", y)
	if p.Board[x][y] != 0 {
		p.Board[x][y] = 255
		p.Hits += 1
		return true, nil
	}
	return false, nil
}

func (p *PlayerBoard) IsDefeated() bool {
	return p.Hits == totalHitsRequired
}
