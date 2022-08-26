package main

import (
	"errors"
	"net"
	"sync"
)

type factory func() (Item, error)

type Item struct {
	conn net.Conn
}

type Response struct {
	i   *Item
	err error
}

var address string

type Pool struct {
	Lock       sync.Mutex
	BusyItem   map[*Item]struct{}
	IdleItem   map[*Item]struct{}
	Ch         chan Response
	Factory    factory
	MaxIdleNum int
	MaxBusyNum int
	ConnNum    int
}

func Init(idle, max int, f factory) *Pool {
	return &Pool{Lock: sync.Mutex{}, BusyItem: map[*Item]struct{}{}, IdleItem: map[*Item]struct{}{}, Ch: make(chan Response), Factory: f, MaxBusyNum: max, MaxIdleNum: idle, ConnNum: 0}
}

func myFactory() (Item, error) {
	conn, err := net.Dial("tpc", address)
	return Item{conn: conn}, err
}

func (p *Pool) Get() {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	if len(p.IdleItem) > 0 {
		for i, _ := range p.IdleItem {
			delete(p.IdleItem, i)
			p.BusyItem[i] = struct{}{}
			p.Ch <- Response{i, nil}
			p.ConnNum++
		}
	} else {
		if p.ConnNum == p.MaxBusyNum {
			p.Ch <- Response{nil, errors.New("too much connect")}
		} else {
			i, err := myFactory()
			p.ConnNum++
			p.Ch <- Response{&i, err}
		}
	}
}

func (p *Pool) GiveBack(i *Item) {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	delete(p.BusyItem, i)
	if len(p.IdleItem) == p.MaxIdleNum {
		i.conn.Close()
	} else {
		p.IdleItem[i] = struct{}{}
	}
	p.ConnNum--
}

func (p *Pool) Release() {
	for i, _ := range p.BusyItem {
		i.conn.Close()
	}
	for i, _ := range p.IdleItem {
		i.conn.Close()
	}
}
