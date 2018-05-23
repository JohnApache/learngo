package main

import (
	"testing"

	"gopkg.in/ffmt.v1"
)

func TestNewPiecesBoard(t *testing.T) {
	NewPiecesBoard(8, 8)
}

func TestHoldPiece(t *testing.T) {
	pb2 := NewPiecesBoard(4, 4)
	p1 := NewPiece(2, 2)
	pb2.HoldPiece(&p1)
	ffmt.Puts(pb2.Board)
	pb2.ClearUse()
	ffmt.Puts(pb2.Board)
}
