/*
This shows what happens when you have multiple go routines subscribing to the same subject controlled by a queue. Only one count subscriber will
get the message, so the message is only processed once.  The shutdown message effects both of the go routines.
*/
package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/nats-io/nats.go"
)

func handleMsgs(procId int, procCh chan string, nc *nats.Conn) {

	// Use a WaitGroup to wait for a message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe

	if _, err := nc.Subscribe("IMAGES.shutdown", func(m *nats.Msg) {
		wg.Done()
	}); err != nil {
		log.Fatal(err)
	}

	// Subscribe to count
	var i int
	if _, err := nc.Subscribe("IMAGES.count", func(m *nats.Msg) {
		i++
		fmt.Printf("proc id: %v count: %v\n", procId, strconv.Itoa(i))
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for a message to come in
	wg.Wait()
	procCh <- "shutdown"
}

func main() {

	procCH := make(chan string)
	var (
		status   string
		natsConn *nats.Conn
	)

	var (
		err error
	)

	if natsConn, err = nats.Connect("nats-dev-1.savup.com", nats.UserCredentials("/Users/syacko/workspace/styh-dev/src/albert/.keys/savup.creds")); err != nil {
		log.Println("Connection to NATS failed")
		log.Fatalln(err.Error())
		// TODO Add Error Handling
	}
	if err != nil {
		log.Fatal(err)
	}
	defer natsConn.Close()

	for i := 1; i < 11; i++ {
		go handleMsgs(i, procCH, natsConn)
	}
	select {
	case status = <-procCH:
		fmt.Printf("Creating-Multi-GoRoutines-with-Channel has %v\n", status)
	}

	fmt.Println("ALL DONE!")
}
