package main

import (
	"testing"
)

func TestA(t *testing.T) {
	f1 := SpiderFile{}
	f1.FileName = "a.txt"
	//	f1.Extend = ".txt"
	f1.Path = "b/c/"
	f1.Content = []byte("bc")
	f1.Write()
}
