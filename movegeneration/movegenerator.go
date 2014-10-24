package movegeneration

import (
	"github.com/peterellisjones/gochess/board"
	"github.com/peterellisjones/gochess/move"
	"github.com/peterellisjones/gochess/side"
)

// Generator generates moves
type Generator struct {
	board *board.Board
}

// New returns a new move generator
func New(bd *board.Board) *Generator {
	return &Generator{
		board: bd,
	}
}

func (gen *Generator) ForEachMove(side side.Side, fn func(move.Move) bool) {
	if gen.ForEachCastle(side, fn) {
		return
	}
	if gen.ForEachQueenMove(side, fn) {
		return
	}
	if gen.ForEachRookMove(side, fn) {
		return
	}
	if gen.ForEachBishopMove(side, fn) {
		return
	}
	if gen.ForEachKnightMove(side, fn) {
		return
	}
	if gen.ForEachKingMove(side, fn) {
		return
	}
	if gen.ForEachPawnMove(side, fn) {
		return
	}
}

func (gen *Generator) AllMoves(side side.Side) []move.Move {
	moveList := []move.Move{}
	gen.ForEachMove(side, func(mv move.Move) bool {
		moveList = append(moveList, mv)
		return false
	})
	return moveList
}

func (gen *Generator) BishopMoves(side side.Side) []move.Move {
	moveList := []move.Move{}
	gen.ForEachBishopMove(side, func(mv move.Move) bool {
		moveList = append(moveList, mv)
		return false
	})
	return moveList
}

func (gen *Generator) KnightMoves(side side.Side) []move.Move {
	moveList := []move.Move{}
	gen.ForEachKnightMove(side, func(mv move.Move) bool {
		moveList = append(moveList, mv)
		return false
	})
	return moveList
}
func (gen *Generator) PawnMoves(side side.Side) []move.Move {
	moveList := []move.Move{}
	gen.ForEachPawnMove(side, func(mv move.Move) bool {
		moveList = append(moveList, mv)
		return false
	})
	return moveList
}
func (gen *Generator) RookMoves(side side.Side) []move.Move {
	moveList := []move.Move{}
	gen.ForEachRookMove(side, func(mv move.Move) bool {
		moveList = append(moveList, mv)
		return false
	})
	return moveList
}
func (gen *Generator) QueenMoves(side side.Side) []move.Move {
	moveList := []move.Move{}
	gen.ForEachQueenMove(side, func(mv move.Move) bool {
		moveList = append(moveList, mv)
		return false
	})
	return moveList
}
func (gen *Generator) KingMoves(side side.Side) []move.Move {
	moveList := []move.Move{}
	gen.ForEachKingMove(side, func(mv move.Move) bool {
		moveList = append(moveList, mv)
		return false
	})
	return moveList
}
func (gen *Generator) Castles(side side.Side) []move.Move {
	moveList := []move.Move{}
	gen.ForEachCastle(side, func(mv move.Move) bool {
		moveList = append(moveList, mv)
		return false
	})
	return moveList
}
