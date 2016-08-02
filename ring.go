package main

import (
	"fmt"
	"time"
)

func main() {
	sendChan := make(chan int)
	done := make(chan bool)
	firstProcess := Process{
		IsFirst: true,
		SendMsg: sendChan,
		Done:    done,
	}
	for i := 1; i < 10000; i++ {
		rcvChan := sendChan
		sendChan = make(chan int)
		process := Process{
			RcvMsg:  rcvChan,
			SendMsg: sendChan,
		}
		go process.Loop()
	}

	firstProcess.RcvMsg = sendChan
	go firstProcess.Loop()

	start := time.Now()
	sendChan <- 10000

	// will wait for the loop to be done
	<-done
	elapsed := time.Since(start)
	fmt.Printf("Loop took %s\n", elapsed)

}

type Process struct {
	RcvMsg  chan int
	SendMsg chan int
	Done    chan bool
	IsFirst bool
}

func (p *Process) Loop() {
	for {
		msg := <-p.RcvMsg
		if p.IsFirst {
			msg--
			if msg < 0 {
				fmt.Println("Finished! Exiting...")
				p.Done <- true
			} else {
				fmt.Println("Finished a loop, starting on", msg)
			}
		}
		p.SendMsg <- msg
	}
}
