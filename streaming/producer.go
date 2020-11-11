package streaming

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

// StartProducer starts the streaming producer
func StartProducer(url, message string) {
	nc, err := nats.Connect(url)
	if err != nil {
		panic(err)
	}
	defer nc.Close()
	sc, err := stan.Connect("test-cluster", "rabbit-test-2", stan.NatsConn(nc))
	if err != nil {
		panic(err)
	}
	defer sc.Close()
	i := 0
	for {
		msg := fmt.Sprintf(message, i)
		sc.Publish("workers", []byte(msg))
		time.Sleep(1 * time.Second)
		i++
	}
}
