package controller

import (
	"fmt"
	"log"
	"os"
	"time"

	"io/ioutil"
	"net/http"
	"encoding/json"

	"go.nanomsg.org/mangos"
	"go.nanomsg.org/mangos/protocol/pub"
	"go.nanomsg.org/mangos/protocol/req"

	// register transports
	_ "go.nanomsg.org/mangos/transport/all"
)

var apiAddress = "http://localhost:8080"
var publisherAddress = "tcp://localhost:40899"
var requestAdress = "tcp://localhost:50899"

type WorkloadStruct struct {
    Workload_Id string `json:"workload_id" binding:"required"`
    Filter string `json:"filter" binding:"required"`
    Workload_Name string `json:"workload_name" binding:"required"`
    Status string `json:"status" binding:"required"`
    Running_Jobs int `json:"running_jobs" binding:"required"`
    Filtered_Images []string `json:"filtered_images" binding:"required"`
}

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
   Message Passing Socket Configuration
*/
func MessagePassingConfig() (pubSock mangos.Socket, reqSock mangos.Socket) {
	// PubSub
	var err error
	if pubSock, err = pub.NewSocket(); err != nil { die("Can't get new pub socket: %s", err) }
	if err = pubSock.Listen(publisherAddress); err != nil { die("Can't listen on pub socket: %s", err.Error()) }

	// ReqRep
	if reqSock, err = req.NewSocket(); err != nil { die("Can't get new req socket: %s", err.Error()) }
	if err = reqSock.Listen(requestAdress); err != nil { die("Can't listen on pub socket: %s", err.Error()) }

	return pubSock, reqSock
}

/*
   Publish
*/
func Publish(sock mangos.Socket) {
	var err error
	d := date()
	log.Printf("Controller: Publishing Date %s\n", d)
	if err = sock.Send([]byte(d)); err != nil { die("Failed publishing: %s", err.Error()) }
	time.Sleep(time.Second * 3)
}

/*
   Request
*/
func Request(sock mangos.Socket) {
	var msg []byte
	var err error
	fmt.Printf("NODE1: SENDING DATE REQUEST %s\n", "DATE")
	sock.Send([]byte("DATE"))
	if msg, err = sock.Recv(); err == nil {
		fmt.Printf("NODE1: RECEIVED DATE %s\n", string(msg))
	}

	//sock.Close()
	time.Sleep(time.Second * 3)
}

/*
   Start
*/
func Start(adminAccess chan string) {

	// API Admin Auth
	token := <-adminAccess
	bearer := "Bearer " + token

	pubSock, reqSock := MessagePassingConfig()

	for {
		// Obtain Workloads List
		req, err := http.NewRequest("GET", apiAddress + "/workloads", nil)
		req.Header.Add("Authorization", bearer)

		client := &http.Client{}
		resp, err := client.Do(req)
    if err != nil { die("Error on request: %s", err.Error()) }
    defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
    if err != nil { die("Error while reading response: %s", err.Error()) }

		// Convert To Struct
		var data []WorkloadStruct
    if err := json.Unmarshal([]byte(string(body)), &data); err != nil { die("Error on data conversion: %s", err.Error()) }
    for _, element := range data {
    	fmt.Println(element.Workload_Name)
    }

		Publish(pubSock)
		go Request(reqSock)

	}
}
