package messaging

import (
	"assessment-go-source-code-muhammad-aditya-reader/internal/model"
	"assessment-go-source-code-muhammad-aditya-reader/internal/usecase"
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type CustomerReader struct {
	Log     *logrus.Logger
	UseCase *usecase.CustomerConsumerUseCase
}

func NewCustomerReader(log *logrus.Logger, useCase *usecase.CustomerConsumerUseCase) *CustomerReader {
	return &CustomerReader{
		Log:     log,
		UseCase: useCase,
	}
}

func (r *CustomerReader) Read(message *kafka.Message) error {
	customerEvent := new(model.CustomerEvent)
	if err := json.Unmarshal(message.Value, customerEvent); err != nil {
		r.Log.WithError(err).Error("error unmarshalling address event")
		return err
	}

	// send to database
	response, err := r.CreateMessageToDatabase(context.Background(), customerEvent)
	if err != nil {
		r.Log.WithError(err).Error("error creating customer consumer")
		return err
	}

	// TODO process event
	r.Log.Infof("Received topic addresses with event: %v from partition %d", customerEvent, message.Partition)
	r.Log.Infof("Customer consumer created: %v", response)
	return nil
}

func (r *CustomerReader) CreateMessageToDatabase(ctx context.Context, customerEvent *model.CustomerEvent) (*model.CustomerConsumerResponse, error) {
	return r.UseCase.Create(ctx, &model.CreateCustomerConsumerRequest{
		NationalId:    customerEvent.NationalId,
		Name:          customerEvent.Name,
		DetailAddress: customerEvent.DetailAddress,
	})
}
