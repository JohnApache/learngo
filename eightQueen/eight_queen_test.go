package main

import (
	"testing"
)

//func TestNewPiecesBoard(t *testing.T) {
//	NewPiecesBoard(8, 8)
//}

//func TestHoldPiece(t *testing.T) {
//	pb2 := NewPiecesBoard(4, 4)
//	p1 := NewPiece(2, 2)
//	pb2.HoldPiece(&p1)
//	ffmt.Puts(pb2.Board)
//	pb2.ClearUse()
//	ffmt.Puts(pb2.Board)
//}

//func TestReStorePieces(t *testing.T) {
//	pb := NewPiecesBoard(4, 4)
//	p1 := NewPiece(2, 2)
//	b := pb.StoreScene()
//	pb.HoldPiece(&p1)
//	b2 := pb.StoreScene()
//	ffmt.Mark("----------------------------")
//	ffmt.Puts(pb.Board)
//	ffmt.Mark("----------------------------")
//	pb.ReStoreScene(b)
//	ffmt.Puts(pb.Board)
//	ffmt.Mark("----------------------------")
//	pb.ReStoreScene(b2)
//	ffmt.Puts(pb.Board)
//	ffmt.Mark("----------------------------")
//}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pb.Run(pb.NextPieces(0), 0)
		pb.ClearUse()
	}
}
