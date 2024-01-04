package sent

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func Sent(email,name string) {
	fmt.Println("Go Rabbit Mq Tutorial")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	fmt.Println("Connected to RabbitMQ!")

	ch, err := conn.Channel()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(q)
	b, _ := json.Marshal(map[string]interface{}{
		"to":      email, // change for tests
		"subject": "Bid Alert!!!!",
		"body":    "New Bid Availabe",
	})

	err = ch.Publish(
		"",
		"TestQue",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         b,
		},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Printf("Succesdsfully Message Published to the queue:%s",name)
}
