package main

import (
	"fmt"
	// Add imports here
)

const (
	// Add Constants here
)

type natsMessage interface {
	natsCall()
}

func natsCall(message natsMessage) {
	message.natsCall()
}

type realSMessage struct {
	owner string
}

func (rm realSMessage) natsCall() {
	fmt.Println("NATS is responding.")
}

type FakeSMessage struct {
	owner string
}

func (fm FakeSMessage) natsCall() {
	fmt.Println("I'm an imposter.")
}

func main() {

	realCalls := []realSMessage{{"Scott"}}
	for _, call := range realCalls {
		call.natsCall()
	}

	fakeCalls := []FakeSMessage{{"Scott"}}
	for _, call := range fakeCalls {
		call.natsCall()
	}

	calls := []natsMessage{realSMessage{"Scott"}, FakeSMessage{"Scott"}}
	for _, call := range calls {
		natsCall(call)
	}
}
