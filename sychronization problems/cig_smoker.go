package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
type Ingredient struct { sync.Mutex }

type Agent struct {
	ingredients [3]*Ingredient 
}

type Smoker struct {
	id int
	i1 *Ingredient
	i2 *Ingredient
}

func randomPause(max int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(max*1000)))
}

func say(action string, id int) {
	fmt.Printf(">%d %s", id, action)
}

func (a Agent) Agent_lock() {
	i1 := rand.Intn(3)
	i2 := rand.Intn(3)
	if(i1==i2) {
		a.Agent_lock()
	}
	a.ingredients[i1].Unlock()
	a.ingredients[i2].Unlock()
}

func (s Smoker) put_scene() {
	say("is awake\n", s.id)
	randomPause(2)

	say("wants to smoke\n", s.id)
	s.i1.Lock()
	s.i2.Lock()

	say("is smoking\n", s.id)
	randomPause(5)

	s.i1.Unlock()
	s.i2.Unlock()
	say("is asleep\n", s.id)
	randomPause(1)

	s.put_scene()
}





func main() {
	//Ingredient 0 is paper
	//Ingredient 1 is tobacco
	//Ingredient 2 is a match box
	ingredients := make([]*Ingredient,3) 

	for i := 0; i < 3; i++ {
		ingredients[i] = new(Ingredient)
		//ingredients[i].Unlock()
	}

	smokers := make([]*Smoker, 3)
	// // agents := make([]*Agent,1)
	// // agent := agents[0]
	// for i := 0; i < 3; i++ {
	// 	agent.ingredients[i] = ingredients[i]
		
	// }


	smokers[0] = &Smoker{ 
			id: 0, i1: ingredients[2],i2: ingredients[1]}
	go smokers[0].put_scene()

	smokers[1] = &Smoker{ 
			id: 1, i1: ingredients[0],i2: ingredients[2]}
	go smokers[1].put_scene()

	smokers[2] = &Smoker{ 
			id: 2, i1: ingredients[1],i2: ingredients[0]}
	go smokers[2].put_scene()

	// //agent.Agent_lock()
	// for i := 0; i < 3; i++ {
	// 	//We are assigning each smoker the locks represnting the ingredients he needs to smoke
	// 	smokers[i] = &Smoker{ 
	// 		id: i, i1: ingredients[(i-1)%3],i2: ingredients[(i+1)%3]}
	// 	go smokers[i].put_scene()
	// }
	//wg.Wait()
	endless := make(chan int)
	<-endless

}