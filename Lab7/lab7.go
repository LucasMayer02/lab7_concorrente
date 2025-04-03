package main

import (
	"fmt"
	"math/rand"
	"time"
)

func exec(maxSleepMs int) int {
	sleep := rand.Intn(maxSleepMs) // tempo aleatório entre 0 e maxSleepMs
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	return sleep
}

// Função que retorna imediatamente um canal sendo preenchido por 1000 execs
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

	// Consumir 1000 valores, independente do canal
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
		}
	}

	fmt.Println("Soma total:", total)
}
