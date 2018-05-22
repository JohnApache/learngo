package main

import (
	"testing"

	"gopkg.in/ffmt.v1"
)

func TestNewPiecesBoard(t *testing.T) {
	pb := NewPiecesBoard(8, 8)
	ffmt.Puts(pb)
}
