package main

import (
	"time"

	"github.com/nats-io/nats"
	"github.com/reconquest/pkg/log"
)

type Request struct {
	ID int
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

	i := 0
	for {
		request := Request{ID: i}
		log.Infof(nil, "sending request, request_id: %d", request.ID)
		chanForSend <- &request
		time.Sleep(time.Second * 1)
		i = i + 1
	}
}
