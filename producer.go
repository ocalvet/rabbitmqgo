package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	config := amqp.Config{}
	fmt.Print(config)
}
