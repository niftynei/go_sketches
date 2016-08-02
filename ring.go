package main

import "fmt"

// TODO: create a ring of go processes that listen for messages
// when they receive a message, they call the next process in the ring
// todo: time how long it takes to send M messages around N processes
func main() {
	sendChan := make(chan int)
	done := make(chan bool)
	firstProcess := Process{
		IsFirst: true,
		SendMsg: sendChan,
		Done:    done,
	}
	for i := 1; i < 5; i++ {
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
	sendChan <- 100

	// will wait for the loop to be done
	<-done
}

type Process struct {
	RcvMsg  chan int
	SendMsg chan int
	Done    chan bool
	IsFirst bool
}

func (p *Process) Loop() {
	for {
		select {
		case msg := <-p.RcvMsg:
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
		default:
			// do nothing
		}
	}
}
