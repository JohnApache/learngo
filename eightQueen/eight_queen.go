package main

import (
	"fmt"
)

var WholeCount int
var pb = NewPiecesBoard(1, 1)

func main() {
	//	八皇后问题是一个以国际象棋为背景的问题：如何能够在 8×8 的国际象棋棋盘上放置八个皇后，
	//	使得任何一个皇后都无法直接吃掉其他的皇后？为了达到此目的，任两个皇后都不能处于同一条横行、
	//	纵行或斜线上。八皇后问题可以推广为更一般的n皇后摆放问题：
	//	这时棋盘的大小变为n1×n1，而皇后个数也变成n2。而且仅当 n2 = 1 或 n1 ≥ 4 时问题有解。
	//	pb := NewPiecesBoard(8, 8)
	pb.Run(pb.NextPieces(0), 0)
	fmt.Println(WholeCount)

}
func (pb *PiecesBoard) Run(ps []Piece, c int) {
	c = c + 1
	if c > pb.maxX || c > pb.maxY {
		WholeCount++
		return
	}
	b := pb.StoreScene()
	for _, v := range ps {
		if v.IsUse == true {
			pb.ReStoreScene(b)
			continue
		}
		pb.HoldPiece(&v)
		nps := pb.NextPieces(c)
		pb.Run(nps, c)
		pb.ReStoreScene(b)
	}
}

type PiecesBoard struct {
	Board [][]Piece
	maxX  int
	maxY  int
}

type Piece struct {
	X     int
	Y     int
	IsUse bool
}

func NewPiece(x, y int) Piece {
	return Piece{
		X:     x,
		Y:     y,
		IsUse: false,
	}
}

func NewPiecesBoard(mx, my int) *PiecesBoard {
	pb := PiecesBoard{
		maxX: mx,
		maxY: my,
	}
	for i := 0; i < mx; i++ {
		HorPieces := []Piece{}
		for j := 0; j < my; j++ {
			HorPieces = append(HorPieces, NewPiece(j, i))
		}
		pb.Board = append(pb.Board, HorPieces)
	}
	return &pb
}

func (pb *PiecesBoard) ClearUse() {
	pb.WalkPieces(func(p *Piece) {
		p.IsUse = false
	})
}

func (pb *PiecesBoard) WalkPieces(f func(p *Piece)) {
	for i, _ := range pb.Board {
		for j, _ := range pb.Board[i] {
			f(&pb.Board[i][j])
		}
	}
}
func (pb *PiecesBoard) RestPieces() []Piece {
	ps := []Piece{}
	pb.WalkPieces(func(p *Piece) {
		if !p.IsUse {
			ps = append(ps, *p)
		}
	})
	return ps
}

func (pb *PiecesBoard) NextPieces(i int) []Piece {
	if i >= len(pb.Board) {
		return nil
	}
	return pb.Board[i]
}

//store current scene
func (pb *PiecesBoard) StoreScene() []bool {
	b := []bool{}
	pb.WalkPieces(func(p *Piece) {
		b = append(b, p.IsUse)
	})
	return b
}

//restore scene
func (pb *PiecesBoard) ReStoreScene(b []bool) {
	i := 0
	pb.WalkPieces(func(p *Piece) {
		p.IsUse = b[i]
		i++
	})
}

//holdAll
func (pb *PiecesBoard) HoldPiece(p *Piece) {
	pb.holdHorPiece(p.Y)
	pb.holdVerPiece(p.X)
	pb.holdRsPiece(p.X, p.Y)
	pb.holdLsPiece(p.X, p.Y)
}

//horizontal hold
func (pb *PiecesBoard) holdHorPiece(y int) {
	hp := pb.Board[y]
	for i, _ := range hp {
		pb.Board[y][i].IsUse = true
	}
}

//vertical hold
func (pb *PiecesBoard) holdVerPiece(x int) {
	for _, v := range pb.Board {
		v[x].IsUse = true
	}
}

//right Slant
func (pb *PiecesBoard) holdRsPiece(x, y int) {
	tempX := x - 1
	tempY := y - 1
	for {
		if x >= pb.maxX || y >= pb.maxY {
			break
		}
		pb.Board[y][x].IsUse = true
		x++
		y++
	}
	for {
		if tempX < 0 || tempY < 0 {
			break
		}
		pb.Board[tempY][tempX].IsUse = true
		tempX--
		tempY--
	}
}

//left Slant
func (pb *PiecesBoard) holdLsPiece(x, y int) {
	tempX := x - 1
	tempY := y + 1
	for {
		if x >= pb.maxX || y < 0 {
			break
		}
		pb.Board[y][x].IsUse = true
		x++
		y--
	}
	for {
		if tempX < 0 || tempY >= pb.maxY {
			break
		}
		pb.Board[tempY][tempX].IsUse = true
		tempX--
		tempY++
	}
}
