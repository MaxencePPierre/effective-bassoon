package main

import (
	"encoding/json"
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/rs/zerolog/log"
)

const (
	localhost string = "127.0.0.1"
	port      string = "4150"

	Topic string = "effective_bassoon"
)

type Message struct {
	Name      string
	Content   string
	Timestamp int64
}

func main() {
	nsqConf := nsq.NewConfig()

	producer, err := nsq.NewProducer(localhost+":"+port, nsqConf)
	if err != nil {
		log.Fatal().Err(err).Msg("NSQ Producer could not connect")
	}

	msg := Message{
		Name:      "Message Name",
		Content:   "Message Content",
		Timestamp: time.Now().Unix(),
	}

	payload, err := json.Marshal(msg)
	if err != nil {
		log.Fatal().Err(err).Msg("Message marshalling error")
	}

	err = producer.Publish(Topic, payload)
	if err != nil {
		log.Fatal().Err(err).Msg("Publish message error")
	}
}