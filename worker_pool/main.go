package main

import (
	"fmt"
	"strconv"
	"sync"
)

var (
	size    = 100
	clients = make(chan Client, size)
	data    = make(chan Data, size)
)

type Client struct {
	id   int
	name string
}

type Data struct {
	client      Client
	jobResponse string
}

func createClients(n int) {
	for i := 0; i < n; i++ {
		s := strconv.Itoa(i)
		c := Client{i, "Name:" + s}
		clients <- c
	}
	close(clients)
}

func worker(w *sync.WaitGroup) {
	for c := range clients {
		data <- Data{c, c.name + " -> Done"}
	}
	w.Done()
}

// the number of Worker Pools we will provide to run the Goroutines
func makeWP(n int) {
	var w sync.WaitGroup
	for i := 0; i < n; i++ {
		w.Add(1)
		go worker(&w)
	}
	w.Wait()
	close(data)
}

func main() {
	nWorkers := 4 // The Number of Worker Pools to be used
	nJobs := 10   // The Number of Clients
	go createClients(nJobs)
	finished := make(chan interface{})
	go func() {
		for d := range data {
			fmt.Println("Client: ", d.client.id, " || ", d.client.name, " || Job: ", d.jobResponse)
		}
		finished <- true
	}()
	makeWP(nWorkers)
	<-finished
}
