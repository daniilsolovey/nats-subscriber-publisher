package main

import (
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

	log.Info("connected to nats and ready to receive messages")

	chanForReceive := make(chan *Request)
	encodedConn.BindRecvChan("request_subject", chanForReceive)

	for {
		request := <-chanForReceive
		log.Infof(nil, "received request, request_id: %d", request.ID)
	}
}
