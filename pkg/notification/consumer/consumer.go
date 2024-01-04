package main

import (
	"fmt"
	"log"

	"github.com/sreerag_v/BidFlow/pkg/notification/mailer"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer Application")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	fmt.Println("Consumer Connected to the Application")

	ch, err := conn.Channel()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	mesg, err := ch.Consume(
		"TestQue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range mesg {
			// if d.Body == nil {
			// 	panic("Nothing Found")
			// }
			// TODO: implement validators
			mailer.SendMail(d.Body)
			log.Printf("Mail sent to: %s", d.Body)
			d.Ack(false)
		}
	}()

	fmt.Println("Sucessfully Coneted To RabbitMq Instance")
	fmt.Println("[*]- waiting for messages")
	<-forever
}
