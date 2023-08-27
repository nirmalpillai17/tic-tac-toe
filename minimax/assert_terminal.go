package minimax

import "github.com/nirmalpillai17/tic-tac-toe/game"

func findMatch(arr [GameSize]game.Player) bool {
	return true
}

func isTerminal(s game.State) bool {
	rows := s.CreateMatrix()
	cols := game.Matrix{}
	dr := [GameSize]game.Player{}  // Slice containing right diagonal
	dl := [GameSize]game.Player{}  // Slice containing left diagonal

	for i, row := range rows {
		dr[i] = row[i]
		dl[i] = row[(GameSize-1)-i]
		for j, c := range row {
			cols[j][i] = c
		}
	}
	return true
}
