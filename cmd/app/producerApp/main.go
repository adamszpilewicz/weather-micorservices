package main

import (
	"encoding/json"
	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
	"log"
	"math/rand"
	"os"
	"rabbit/conf"
	db2 "rabbit/db"
	"rabbit/producer"
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

	// ----- producer part -----
	producer, err := producer.NewProducer(conf.RabbitAddress)
	if err != nil {
		log.Fatal(err)
	}

	err = producer.CreateQueue(conf.QueueName)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 1000; i++ {
		age := rand.Int63n(200)
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		msgJson, err := json.Marshal(db2.Payload{Name: "Adam", Age: int(age), Date: time.Now(), Unix: time.Now().Unix()})
		if err != nil {
			log.Fatal(err)
		}
		payload := string(msgJson)
		err = producer.SendMessage(conf.QueueName, payload)
		if err != nil {
			log.Fatal(err)
		}
	}

}
