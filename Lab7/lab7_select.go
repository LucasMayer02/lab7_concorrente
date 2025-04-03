package main

import (
	"fmt"
	"math/rand"
	"time"
)

func exec(maxSleepMs int) int {
	sleep := rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	return sleep
}

func startProducer(maxSleepMs int) <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 1000; i++ {
			out <- exec(maxSleepMs)
		}
		close(out)
	}()

	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())

	c1 := startProducer(10)
	c2 := startProducer(10)

	total := 0
	received := 0

	for received < 1000 {
		select {
		case v, ok := <-c1:
			if ok {
				total += v
				received++
			}
		case v, ok := <-c2:
			if ok {
				total += v
				received++
			}
		default:
			// Canal ainda não pronto? Evita bloqueio.
			time.Sleep(time.Millisecond) // yield
		}
	}

	fmt.Println("Soma total com select otimizado:", total)
}
