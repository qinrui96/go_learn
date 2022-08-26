package main

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	m       sync.Mutex
	res     chan io.Closer
	factory func() (io.Closer, error)
	closed  bool
}

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	return &Pool{factory: fn, res: make(chan io.Closer, size)}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.res:
		log.Println("Acquire:shared")
		if !ok {
			return nil, errors.New("closed")
		}
		return r, nil
	default:
		log.Println("Acquire:new")
		return p.factory()

	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return
	}
	p.closed = true

	close(p.res)

	for r := range p.res {
		r.Close()
	}
}
