package n_queens

func cNN(l uint8) uint64 {
	var c = newChessboardNN(l)
	var sum uint64
	c.nQueens(0, l, &sum)
	return sum
}

func newChessboardNN(i uint8) chessboardNN {
	m := i >> 3
	if i&0x7 != 0 {
		m++
	}
	return chessboardNN{
		m: m,
		c: make([]chessboard88, m*m),
	}
}

type chessboardNN struct {
	c []chessboard88
	m uint8
}

func (c chessboardNN) Point(x, y uint8) bool {
	px := x >> 3
	cx := x & 0x7
	py := y >> 3
	cy := y & 0x7
	index := px + py*c.m
	return c.c[index].Point(cx, cy)
}

func (c chessboardNN) SetPointTrue(x, y uint8) chessboardNN {
	px := x >> 3
	cx := x & 0x7
	py := y >> 3
	cy := y & 0x7
	index := px + py*c.m
	c.c[index] = c.c[index].SetPointTrue(cx, cy)
	return c
}

func (c chessboardNN) Clone() chessboardNN {
	r := make([]chessboard88, len(c.c))
	copy(r, c.c)
	return chessboardNN{
		m: c.m,
		c: r,
	}
}

func (c chessboardNN) nQueens(y uint8, l uint8, sum *uint64) {
	for x := uint8(0); x != l; x++ {
		if c.Point(x, y) {
			continue
		}
		yy := y + 1
		if yy == l {
			(*sum)++
			return
		}
		c.Clone().
			setPointTrue(x, y, l).
			nQueens(yy, l, sum)
	}
}

func (c chessboardNN) setPointTrue(x, y uint8, l uint8) chessboardNN {
	for i, j := x+1, y+1; i != l && j != l; {
		c.SetPointTrue(i, j)
		i++
		j++
	}

	for i, j := x-1, y+1; i != 0xFF && j != l; {
		c.SetPointTrue(i, j)
		i--
		j++
	}

	for j := y + 1; j != l; j++ {
		c.SetPointTrue(x, j)
	}
	return c
}
