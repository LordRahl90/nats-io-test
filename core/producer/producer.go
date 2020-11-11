package producer

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

// Start - entrypoint to start the producer
func Start(ctx context.Context, url, message string) {
	nc, err := nats.Connect(url)
	if err != nil {
		panic(err)
	}
	defer func() {
		nc.Close()
	}()

	i := 0
	for {
		fmt.Printf("Sent %d\n", i)
		if err := nc.Publish("messages", []byte(fmt.Sprintf(message, i))); err != nil {
			panic(err)
		}
		i++
		time.Sleep(time.Second * 1)
	}

}
