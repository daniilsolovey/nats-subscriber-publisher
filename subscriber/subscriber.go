package main

import (
	"sync"

	"github.com/nats-io/nats"
	"github.com/reconquest/pkg/log"
)

type Request struct {
	ID int
}

type Request_3 struct {
	Name string
}

func main() {
	natsConn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	encodedConn, err := nats.NewEncodedConn(natsConn, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	defer encodedConn.Close()

	log.Info("connected to nats and ready to receive messages")

	var wg sync.WaitGroup
	wg.Add(1)
	chanForReceive_1 := make(chan *Request)
	encodedConn.BindRecvChan("request_subject", chanForReceive_1)
	go func() {
		for {
			request := <-chanForReceive_1
			log.Infof(nil, "received request for subscriber_1, request_id: %d", request.ID)
		}

	}()

	wg.Add(1)
	chanForReceive_2 := make(chan *Request)
	encodedConn.BindRecvChan("request_subject", chanForReceive_2)
	go func() {
		for {
			request := <-chanForReceive_2
			log.Infof(nil, "received request for subscriber_2, request_id: %d", request.ID)
		}

	}()

	wg.Add(1)
	chanForReceive_3 := make(chan *Request_3)
	encodedConn.BindRecvChan("request_subject_3", chanForReceive_3)

	go func() {
		for {
			request := <-chanForReceive_3
			log.Infof(nil, "received request for subscriber_2, request_id: %s", request.Name)
		}

	}()

	wg.Wait()
}
