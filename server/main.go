package main

import (
	"context"
	"crypto/tls"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func main() {
	mechanism, err := scram.Mechanism(scram.SHA512,
		"{{ UPSTASH_KAFKA_USERNAME }}", "{{ UPSTASH_KAFKA_PASSWORD }}")
	if err != nil {
		log.Fatalln(err)
	}

	dialer := &kafka.Dialer{
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"{{ BOOTSTRAP_ENDPOINT }}"},
		GroupID: "{{ GROUP_NAME }}",
		Topic:   "{{ TOPIC_NAME }}",
		Dialer:  dialer,
	})
	defer r.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	m, err := r.ReadMessage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", m)
}
