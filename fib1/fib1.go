package fib1

import (
	"strconv"

	//	ffmt "gopkg.in/ffmt.v1"
)

//给定一个数字组成的字符串，判断字符串是否可以分割为类斐波那契数列A[n]=A[n-1]+A[n-2] ( for n>=2)
//输入：字符串(纯数字)
//输出：True  or False
//如:1123581321 -> [1,1,2,3,5,8,13,21]  输出：True
//如:100101201302 -> [100,101,201,302] 输出：True
//如:112211221122  输出:False

func Run() {
	t := newT([]byte("1123581321"))
	t.TryAll()
	//	{
	//		t := newT([]byte("1123581321"))
	//		ffmt.Mark(t.TryAll())
	//	}
	//	{
	//		t := newT([]byte("100101201302"))
	//		ffmt.Mark(t.TryAll())
	//	}
	//	{
	//		t := newT([]byte("112211221122"))
	//		ffmt.Mark(t.TryAll())
	//	}
}

type T struct {
	buf []byte
	off int
	ch  byte
}

func newT(data []byte) *T {
	t := &T{
		buf: data,
	}
	t.next()
	return t
}

func (t *T) TryAll() []int {
	max := len(t.buf) / 3
	for i := 0; i != max; i++ {
		for j := i; j != max; j++ {
			t.Reset()
			b := t.Try(i+1, j+1)
			if b != nil {
				return b
			}
		}
	}
	return nil
}

func (t *T) Reset() {
	t.off = 0
	t.next()
}

func (t *T) next() bool {
	if len(t.buf) < t.off {
		return false
	} else if len(t.buf) == t.off {
		t.ch = 0
		t.off++
		return true
	}
	t.ch = t.buf[t.off]
	t.off++
	return true
}

func (t *T) scan(x int) int {
	s := 0
	for i := 0; i != x; i++ {
		s *= 10
		s += int(t.ch - '0')
		if !t.next() {
			return -1
		}
	}
	return s
}

func (t *T) Try(x1, x2 int) []int {
	r := []int{}
	s1 := t.scan(x1)
	if s1 == -1 {
		return nil
	}
	s2 := t.scan(x2)
	if s2 == -1 {
		return nil
	}

	r = append(r, s1)
	r = append(r, s2)
	return t.try(r)
}

func (t *T) try(r []int) []int {
	if t.ch == 0 {
		return r
	}
	lr := len(r)
	b := r[lr-1] + r[lr-2]
	bb := []byte(strconv.FormatUint(uint64(b), 10))
	s3 := t.scan(len(bb))
	if s3 != b {
		return nil
	}
	r = append(r, b)
	return t.try(r)
}
