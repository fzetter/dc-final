package controller

import (
	"fmt"
	//"log"
	"os"
	"time"

	"go.nanomsg.org/mangos"
	"go.nanomsg.org/mangos/protocol/pub"

	// register transports
	_ "go.nanomsg.org/mangos/transport/all"
)

var controllerAddress = "tcp://localhost:40899"

/*
   Die
*/
func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

/*
   Date
*/
func date() string {
	return time.Now().Format(time.ANSIC)
}

/*
   Start
*/
func Start() {
	var sock mangos.Socket
	var err error

	if sock, err = pub.NewSocket(); err != nil { die("Can't get new pub socket: %s", err) }

	if err = sock.Listen(controllerAddress); err != nil { die("Can't listen on pub socket: %s", err.Error()) }

	for {
		// Obtain workloads list

		// Could also use sock.RecvMsg to get header
		d := date()
		//log.Printf("Controller: Publishing Date %s\n", d)

		if err = sock.Send([]byte(d)); err != nil { die("Failed publishing: %s", err.Error()) }

		time.Sleep(time.Second * 3)
	}
}
