package n_queens

import "fmt"

func cc(l uint8) uint64 {
	if l <= 8 {
		return c88(l)
	}
	return cNN(l)
}

type Pointer interface {
	Point(x, y uint8) bool
}

func Format(c Pointer, l uint8) string {
	buf := make([]byte, 0, l*(l+1))
	for j := uint8(0); j != l; j++ {
		for i := uint8(0); i != l; i++ {
			if c.Point(i, j) {
				buf = append(buf, 'X')
			} else {
				buf = append(buf, '+')
			}
		}
		buf = append(buf, '\n')
	}
	return string(buf)
}

func FormatPrintln(c Pointer, l uint8) {
	fmt.Println(Format(c, l))
}
