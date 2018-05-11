package mapTree

import (
	"testing"

	"gopkg.in/ffmt.v1"
)

func TestA(t *testing.T) {
	mt := NewMT(3)
	mt.AddNode(1)
	ffmt.Puts(mt)
	mt.AddNode(2)
	ffmt.Puts(mt)
	mt.DeleteNode(1)
	ffmt.Puts(mt)
}
