package main

import (
	"fmt"
	"os"
	"os/signal"
	streaming "rabbit_test/streaming"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	url := "nats://127.0.0.1:4223"

	go streaming.StartProducer(url, "hello world %d")

	go func() {
		time.Sleep(time.Second * 5)
		for i := 1; i <= 5; i++ {
			streaming.StartConsumer(url, done)
		}
	}()

	// ctx := context.Background()

	// go producer.Start(ctx, url, "hello world %d")
	// for i := 1; i <= 5; i++ {
	// 	// time.Sleep(1 * time.Second)
	// 	go consumer.Start(ctx, url)
	// }

	<-sigs
	done <- true
	fmt.Println("All Done")
}
