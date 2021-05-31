package controller

import (
	"fmt"
	//"log"
	"os"
	"time"

	"io/ioutil"
	"net/http"
	"encoding/json"

	"go.nanomsg.org/mangos"
	"go.nanomsg.org/mangos/protocol/pub"

	// register transports
	_ "go.nanomsg.org/mangos/transport/all"
)

var apiAddress = "http://localhost:8080"
var controllerAddress = "tcp://localhost:40899"

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
   Start
*/
func Start(adminAccess chan string) {

	// API Admin Auth
	token := <-adminAccess
	bearer := "Bearer " + token

	var sock mangos.Socket
	var err error

	// Create Socket
	if sock, err = pub.NewSocket(); err != nil { die("Can't get new pub socket: %s", err) }

	// Listen Socket
	if err = sock.Listen(controllerAddress); err != nil { die("Can't listen on pub socket: %s", err.Error()) }

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


		// Could also use sock.RecvMsg to get header
		d := date()
		//log.Printf("Controller: Publishing Date %s\n", d)

		if err = sock.Send([]byte(d)); err != nil { die("Failed publishing: %s", err.Error()) }

		time.Sleep(time.Second * 3)
	}
}
