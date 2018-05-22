package main

func main() {
	//	八皇后问题是一个以国际象棋为背景的问题：如何能够在 8×8 的国际象棋棋盘上放置八个皇后，
	//	使得任何一个皇后都无法直接吃掉其他的皇后？为了达到此目的，任两个皇后都不能处于同一条横行、
	//	纵行或斜线上。八皇后问题可以推广为更一般的n皇后摆放问题：
	//	这时棋盘的大小变为n1×n1，而皇后个数也变成n2。而且仅当 n2 = 1 或 n1 ≥ 4 时问题有解。

}

var (
	maxX = 8
	maxY = 8
)

type PiecesBoard struct {
	Board []Piece
}

type Piece struct {
	X     int
	Y     int
	IsUse bool
}

func (p *Piece) HoldPieces() *[]Piece {
	ps := []Piece{}
	for i:=0;i<maxX
	ps = append(ps, *p)
	return nil
}

func NewPiece(x, y int) Piece {
	return Piece{
		X:     x,
		Y:     y,
		IsUse: false,
	}
}

func NewPiecesBoard(mx, my int) *PiecesBoard {
	pb := PiecesBoard{}
	for i := 0; i < mx; i++ {
		for j := 0; j < my; j++ {
			pb.Board = append(pb.Board, NewPiece(i, j))
		}
	}
	return &pb
}

func (pb *PiecesBoard) ClearUse() {
	for _, v := range pb.Board {
		v.IsUse = false
	}
}

func (pb *PiecesBoard) HoldPiece() {

}
