package consumer

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
)

// Start - reads from the queue and print out.
func Start(ctx context.Context, url string) {
	nc, err := nats.Connect(url)
	if err != nil {
		panic(err)
	}

	// ch := make(chan *nats.Msg, 64)
	// for msg := range ch {
	// 	fmt.Printf("Recieved: %s\n\n", msg.Data)
	// }
	for {
		_, err := nc.QueueSubscribe("messages", "recievers", func(msg *nats.Msg) {
			fmt.Printf("Recieved: %s\n\n", msg.Data)
		})
		if err != nil {
			break
		}
	}
	nc.Close()
	fmt.Println("All done now")
	ctx.Done()
}
