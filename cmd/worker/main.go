package main

import (
	"assessment-go-source-code-muhammad-aditya-reader/internal/config"
	"assessment-go-source-code-muhammad-aditya-reader/internal/delivery/messaging"
	"assessment-go-source-code-muhammad-aditya-reader/internal/repository"
	"assessment-go-source-code-muhammad-aditya-reader/internal/usecase"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/segmentio/kafka-go"
)

func main() {
	viperConfig := config.NewViper()
	logger := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, logger)
	validate := config.NewValidator(viperConfig)
	logger.Info("Starting worker service")

	ctx, cancel := context.WithCancel(context.Background())

	logger.Info("setup customer reader")
	customerRepository := repository.NewCustomerConsumerRepository(logger)
	customerUseCase := usecase.NewCustomerConsumerUseCase(db, logger, validate, customerRepository)
	customerReader := config.NewKafkaReader(viperConfig, logger, "customers")
	customerHandler := messaging.NewCustomerReader(logger, customerUseCase)
	go messaging.ReadTopic(ctx, customerReader, logger, func(message kafka.Message) error { // anonymous function as parameter (tanpa nama)
		return customerHandler.Read(&message)
	})

	logger.Info("Worker is running")

	// Membuat channel untuk menangkap sinyal SIGINT dan SIGTERM
	// Sinyal SIGINT dan SIGTERM akan dikirimkan ke channel terminateSignals
	// Ketika sinyal diterima, maka akan memanggil cancel untuk memberitahu konteks agar berhenti
	// Kemudian menunggu konteks selesai (semua proses selesai)
	// Setelah itu, worker akan berhenti
	// Worker akan berhenti ketika menerima sinyal SIGINT atau SIGTERM
	// atau ketika semua proses selesai
	// atau ketika terjadi error
	// atau ketika terjadi panic
	// atau ketika worker dihentikan secara paksa
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
