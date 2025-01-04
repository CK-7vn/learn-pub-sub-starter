package main

import (
	"fmt"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	connString := "amqp://guest:guest@localhost:5672/"
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	conn, err := amqp.Dial(connString)
	if err != nil {
		fmt.Println("Failed to connect to rabbit")
		os.Exit(1)
	} else {
		fmt.Println("Connection succesful to rabbit!")
	}
	defer conn.Close()
	fmt.Println("Starting Peril server...")
	<-signalChan
	fmt.Println("Shutting down Peril server...")

}
