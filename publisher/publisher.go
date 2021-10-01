package main

import (
	"strconv"
	"sync"
	"time"

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

	log.Info("connected to nats and ready to send messages")

	chanForSend := make(chan *Request)
	encodedConn.BindSendChan("request_subject", chanForSend)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		i := 0
		for {
			request := Request{ID: i}
			log.Infof(nil, "sending request, request_id: %d", request.ID)
			chanForSend <- &request
			time.Sleep(time.Second * 1)
			i = i + 1
		}
	}()

	chanForSend_3 := make(chan *Request_3)
	encodedConn.BindSendChan("request_subject_3", chanForSend_3)

	wg.Add(1)
	go func() {
		i := 0
		for {
			request_3 := Request_3{Name: "testName_xxxx: " + strconv.Itoa(i)}
			log.Infof(nil, "sending request to subscriber_3, request_id: %s", request_3.Name)
			chanForSend_3 <- &request_3
			time.Sleep(time.Second * 1)
			i = i + 1
		}
	}()

	wg.Wait()
}
