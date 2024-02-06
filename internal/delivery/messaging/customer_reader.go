package messaging

import (
	"assessment-go-source-code-muhammad-aditya-reader/internal/model"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type CustomerReader struct {
	Log *logrus.Logger
}

func NewCustomerReader(log *logrus.Logger) *CustomerReader {
	return &CustomerReader{
		Log: log,
	}
}

func (r *CustomerReader) Read(message *kafka.Message) error {
	customerEvent := new(model.CustomerEvent)
	if err := json.Unmarshal(message.Value, customerEvent); err != nil {
		r.Log.WithError(err).Error("error unmarshalling address event")
		return err
	}

	// TODO process event
	r.Log.Infof("Received topic addresses with event: %v from partition %d", customerEvent, message.Partition)
	return nil
}
