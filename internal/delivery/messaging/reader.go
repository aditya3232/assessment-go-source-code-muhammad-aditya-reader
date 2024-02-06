package messaging

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type ReaderHandler func(message kafka.Message) error

func ReadTopic(ctx context.Context, reader *kafka.Reader, log *logrus.Logger, handler ReaderHandler) {
	run := true

	for run {
		select {
		case <-ctx.Done():
			run = false
		default:
			message, err := reader.ReadMessage(ctx)
			if err == nil {
				err := handler(message)
				if err != nil {
					log.Errorf("Failed to process message: %v", err)
				} else {
					err = reader.CommitMessages(ctx, message)
					if err != nil {
						log.Fatalf("Failed to commit message: %v", err)
					}
				}
			} else if err != context.Canceled && err != context.DeadlineExceeded {
				log.Warnf("Reader error: %v", err)
			}
		}
	}

	log.Infof("Closing reader for topic: %s", reader.Config().Topic)
	err := reader.Close()
	if err != nil {
		panic(err)
	}
}
