package main

import "github.com/hillfolk/go-kafka-microservice/pkg/events"

func (e events.Event) Process() (*Investment, error) {
	return nil, nil
}
