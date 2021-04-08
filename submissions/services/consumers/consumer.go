package consumers

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type Consumer struct {
	queue string
	handlers []func(data []byte)
}

func NewConsumer(queue string) *Consumer {
	return &Consumer{queue: queue}
}

func (c *Consumer) CallHandlers(body []byte) {
	for _, handler := range c.handlers {
		handler(body)
	}
}

func (c *Consumer) OnReceivedMessage(handler func(data []byte)) {
	c.handlers = append(c.handlers, handler)
}

func (c *Consumer) Init() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	c.handleError(err)
	defer conn.Close()

	ch, err := conn.Channel()
	c.handleError(err)
	defer ch.Close()

	messages, err := ch.Consume(c.queue, "", false, false, false, false, nil)

	forever := make(chan bool)
	go func() {
		for d := range messages {
			c.CallHandlers(d.Body)
		}
	}()

	fmt.Printf(" [*] - waiting for queue %s messages\n", c.queue)
	<-forever
}

func (c *Consumer) handleError(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}
