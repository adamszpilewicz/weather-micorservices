package elas

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	es "github.com/elastic/go-elasticsearch/v8"
	"rabbit/conf"
	"rabbit/db"
)

type EsStorage struct {
	Client *es.Client
}

func NewEsClient(conf conf.Config) (*es.Client, error) {
	client, err := es.NewClient(es.Config{Addresses: conf.EsAddress})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (es EsStorage) Save(payload db.PayloadWeather) error {

	res, err := es.Client.Index("weather", esutil.NewJSONReader(&payload))
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
