// Package main.go
/*
We are showing how to create a server that loops with signaling

RESTRICTIONS:
	None

NOTES:
	Information about Signal:
	https://www.computerhope.com/unix/signals.htm

	GO package: https://pkg.go.dev/os/signal


COPYRIGHT 2022 STY Holdings and/or Scott Yacko
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.

*/
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Add Constants here

// List type's here

var (
// Add Variables here for the file (Remember, they are global)
)

//func main() {
//
//	log.Println("This program will run forever, until you press CTRL-C")
//
//	go serverLoop()
//
//	c := make(chan os.Signal, 1)
//
//	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGINFO)
//
//	go func() {
//		for {
//			select {
//			case sig := <-c:
//				log.Printf("Trapped %q signal", sig)
//				switch sig {
//				case syscall.SIGINT:
//					log.Println("SIGINT Shutdown - You pressed CTRL-C")
//					os.Exit(0)
//				case syscall.SIGTERM:
//					log.Println("SIGTERM Shutdown - You pressed ")
//					os.Exit(0)
//				case syscall.SIGINFO:
//					log.Println("SIGINFO - You pressed CTRL-T")
//					log.Println("==> The server is running.")
//				}
//			case <-quitChannel:
//				return
//			}
//		}
//	}()
//}

func main() {

	log.Println("This program will run forever, until you press CTRL-C")

	go serverLoop()

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

	fmt.Println("\nServer is shutting down!")
}

func serverLoop() {
	for {
		fmt.Printf("%v+\n", time.Now())
		time.Sleep(time.Second)
	}
}
