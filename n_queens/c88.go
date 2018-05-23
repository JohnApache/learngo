package n_queens

func c88(l uint8) uint64 {
	var c chessboard88
	var sum uint64
	c.nQueens(0, l, &sum)
	return sum
}

type chessboard88 uint64

func (c chessboard88) Point(x, y uint8) bool {
	return c&point(x, y) != 0
}

func (c chessboard88) SetPointTrue(x, y uint8) chessboard88 {
	return c | point(x, y)
}

func (c chessboard88) nQueens(y uint8, l uint8, sum *uint64) {
	for x := uint8(0); x != l; x++ {
		if c.Point(x, y) {
			continue
		}
		yy := y + 1
		if yy == l {
			(*sum)++
			return
		}
		c.setPointTrue(x, y, l).
			nQueens(yy, l, sum)
	}
}

func (c chessboard88) setPointTrue(x, y uint8, l uint8) chessboard88 {
	for i, j := x+1, y+1; i != l && j != l; {
		c = c.SetPointTrue(i, j)
		i++
		j++
	}

	for i, j := x-1, y+1; i != 0xFF && j != l; {
		c = c.SetPointTrue(i, j)
		i--
		j++
	}

	for j := y + 1; j != l; j++ {
		c = c.SetPointTrue(x, j)
	}
	return c
}

func point(x, y uint8) chessboard88 {
	return chessboard88(1 << (x + y*8))
}
