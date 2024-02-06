package main

import (
	"assessment-go-source-code-muhammad-aditya-reader/internal/config"
	"assessment-go-source-code-muhammad-aditya-reader/internal/delivery/messaging"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/segmentio/kafka-go"
)

func main() {
	viperConfig := config.NewViper()
	logger := config.NewLogger(viperConfig)
	logger.Info("Starting worker service")

	ctx, cancel := context.WithCancel(context.Background())

	logger.Info("setup customer reader")
	customerReader := config.NewKafkaReader(viperConfig, logger, "customers")
	customerHandler := messaging.NewCustomerReader(logger)
	// Membungkus metode Read dari CustomerReader ke dalam sebuah fungsi yang sesuai dengan ReaderHandler
	handlerFunc := func(message kafka.Message) error {
		return customerHandler.Read(&message)
	}

	go messaging.ReadTopic(ctx, customerReader, logger, handlerFunc)

	logger.Info("Worker is running")

	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGTERM)

	// Menunggu sinyal SIGINT atau SIGTERM
	<-terminateSignals
	logger.Info("Got one of stop signals, shutting down worker gracefully")

	// Memanggil cancel untuk memberitahu konteks agar berhenti
	cancel()

	// Menunggu konteks selesai (semua proses selesai)
	<-ctx.Done()
}
