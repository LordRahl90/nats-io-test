package streaming

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

// StartConsumer start the consumer service
func StartConsumer(url string, done chan bool) {
	nc, err := nats.Connect(url)
	if err != nil {
		panic(err)
	}
	defer nc.Close()
	sc, err := stan.Connect("test-cluster", "rabbit-test-client", stan.NatsConn(nc))
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	_, err = sc.QueueSubscribe("workers", "workers", func(m *stan.Msg) {
		t := time.Unix(m.Timestamp, 0)
		fmt.Printf("Recieved: %s\t on %v\n", m.Data, t.Year())
	}, stan.DurableName("durable-name"))

	if err != nil {
		panic(err)
	}

	<-done
	fmt.Printf("All consumers done\n")
}
