package main

import (
	"context"
	"encoding/json"
	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
	"log"
	"os"
	"rabbit/conf"
	"rabbit/consumer"
	db2 "rabbit/db"
	"time"
)

func main() {
	// ----- config part ------
	vi := viper.New()
	vi.SetConfigFile("./conf/app/conf.yaml")
	vi.ReadInConfig()

	logger := slog.New(slog.NewJSONHandler(os.Stdout))
	slog.SetDefault(logger)
	conf := conf.Config{
		RabbitAddress: vi.GetString("RABBIT_ADDRESS"),
		PgAddress:     vi.GetString("POSTGRES_ADDRESS"),
		StorageType:   vi.GetString("STORAGE_TYPE"),
		QueueName:     vi.GetString("QUEUE_NAME"),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Hour)
	defer cancel()

	// ----- db part ------
	db, err := db2.NewStorage(conf)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// ----- consumer part -----
	consumer, err := consumer.NewConsumer(conf.RabbitAddress)
	if err != nil {
		log.Fatal(err)
	}
	msgs, err := consumer.ConsumeMessage(conf.QueueName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Reading messages")
	log.Println("----------------")

	go func() {
		for msg := range msgs {
			var payloadUnmarshalled db2.Payload
			err := json.Unmarshal(msg.Body, &payloadUnmarshalled)
			if err != nil {
				log.Fatal(err)
			}
			err = db.Save(payloadUnmarshalled)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	<-ctx.Done()

}
