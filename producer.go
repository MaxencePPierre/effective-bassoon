package main

import (
	"encoding/json"

	"github.com/MaxencePPierre/effective-bassoon/message"
	"github.com/nsqio/go-nsq"
	"github.com/rs/zerolog/log"
)

const (
	localhost string = "127.0.0.1"
	port      string = "4150"
)

func main() {
	nsqConf := nsq.NewConfig()

	producer, err := nsq.NewProducer(localhost+":"+port, nsqConf)
	if err != nil {
		log.Fatal().Err(err).Msg("NSQ Producer could not connect")
	}

	msg := message.New("Message Name", "Message Content")
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Fatal().Err(err).Msg("Message marshalling error")
	}

	err = producer.Publish(message.Topic, payload)
	if err != nil {
		log.Fatal().Err(err).Msg("Publish message error")
	}
}
