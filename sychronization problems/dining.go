package main

import (

	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Fork struct{sync.Mutex }

type Philosopher struct {
	id int
	leftFork, rightFork *Fork
}

func (p Philosopher) dine() {
	say("thinking", p.id)
	randomPause(2)

	say("hungry", p.id)
	p.leftFork.Lock()
	p.rightFork.Lock()

	say("eating", p.id)
	randomPause(5)

	p.rightFork.Unlock()
	p.leftFork.Unlock()

	p.dine()

}

func randomPause (max int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(max*1000)))
}

func say(action string, id int) {
	fmt.Printf(">%d is %s\n", id, action)
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	count := 5

	forks := make([]*Fork, count)
	for i := 0; i < count; i++ {
		forks[i] = new(Fork)
	}

	philosophers := make([]*Philosopher, count)

	for i := 0; i < count; i++ {
		philosophers[i] = &Philosopher{
			id: i, leftFork: forks[i], rightFork: forks[(i+1)%count]}
			go philosophers[i].dine()
		
	}

	endless := make(chan int)
	<- endless
	

}