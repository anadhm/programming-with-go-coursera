package main

/*
	Implementation of the dining philosophers problem, with the following modifications:

	1.There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
    2.Each philosopher should eat only 3 times.
    3.The philosophers pick up the chopsticks in any order, not lowest-numbered first.
    4.In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
    5.The host allows no more than 2 philosophers to eat concurrently.
    6.Each philosopher is numbered, 1 through 5.
    7.When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
	8.When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/
import (
	"fmt"
	rand "math/rand"
	"sync"
	"time"
)

var can sync.Mutex
var wg sync.WaitGroup

type philosopher struct {
	number int
	lChop  chopstick
	rChop  chopstick
}
type chopstick struct {
	number int
	mux    sync.Mutex
}

/*
	Routines host and eat communicate via a channel called ask.
	Eat() sends "start" on the channel, and when there are more than 2 requests the send blocks, so capacity is enforced by the channel's buffer.
	If channel is at full capacity, host consumes 2 requests so it is empty again and the routines that made the requests can run.
	Which chopstick is selected to be locked first is chosen randomly.
	If the left is chosen first, then right is unlocked first when eat() is done, and vice versa.
*/
func host(ask chan string) {
	//defer wg.Done()
	for {
		if len(ask) == 2 {
			<-ask
			<-ask
		}
	}
}
func (p philosopher) eat(ask chan string) {
	defer wg.Done()
	i := 0
	for i < 3 {
		ask <- "start"

		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(2)
		if r == 0 {
			p.lChop.mux.Lock()
			p.rChop.mux.Lock()
		} else {
			p.rChop.mux.Lock()
			p.lChop.mux.Lock()
		}
		fmt.Printf("starting to eat %d\n", p.number)

		fmt.Printf("finishing eating %d\n", p.number)
		if r == 0 {
			p.rChop.mux.Unlock()
			p.lChop.mux.Unlock()
		} else {
			p.lChop.mux.Unlock()
			p.rChop.mux.Unlock()
		}
		i++ //finished a round of eating
	}
}
func main() {
	//init
	ask := make(chan string, 2)
	//ok := make(chan bool)
	var philos [5]philosopher
	var chopsticks [5]chopstick
	for i := 0; i < 5; i++ {
		chopsticks[i] = chopstick{i, sync.Mutex{}}
		philos[i] = philosopher{i + 1, chopsticks[i], chopsticks[(i+1)%5]}
	}
	//wg.Add(1)
	go host(ask)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(ask)
	}
	wg.Wait()

}
