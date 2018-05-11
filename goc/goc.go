package goc

import (
	"net/http"
	"sync"
)

var none = struct{}{}

type GRC struct {
	WorkBuf chan func()
	GoMax   chan struct{}
	Wg      sync.WaitGroup
}

func TaskA() {
	http.Get("https://www.boanjie.com")
}

func Goc(mw, mg int) *GRC {
	return &GRC{
		WorkBuf: make(chan func(), mw),
		GoMax:   make(chan struct{}, mg),
	}
}

func (grc *GRC) Push(f func()) {
	grc.Wg.Add(1)
	select {
	case grc.GoMax <- none:
		go grc.Execute(f)
	default:
		grc.WorkBuf <- f
	}
}
func (grc *GRC) Execute(f func()) {
	if f != nil {
		f()
		grc.Wg.Done()
	}
	select {
	case f0 := <-grc.WorkBuf:
		grc.Execute(f0)
	default:
	}
	<-grc.GoMax
}

func (grc *GRC) Join() {
	grc.Wg.Wait()
	return
}
