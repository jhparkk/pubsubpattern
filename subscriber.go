package main

import (
	"context"
	"fmt"
	"time"
)

type Subscriber struct {
	ctx   context.Context
	name  string
	msgCh chan string
}

func NewSubscriber(name string, ctx context.Context) *Subscriber {
	return &Subscriber{
		ctx:   ctx,
		name:  name,
		msgCh: make(chan string),
	}
}

func (s *Subscriber) Subscribe(pub *Publisher) {
	pub.Subscribe(s.msgCh)
}

func (s *Subscriber) Update() {
	loop := 0
	for {
		select {
		case msg := <-s.msgCh:
			fmt.Printf("%d %s got Message : %s %d\n", time.Now().Unix(), s.name, msg, loop)
		case <-s.ctx.Done():
			wg.Done()
			return
		}
		loop++
	}
}
