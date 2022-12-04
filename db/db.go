package db

import (
	"rabbit/conf"
	"strings"
	"time"
)

type Payload struct {
	Name string    `json:"name"`
	Age  int       `json:"age"`
	Date time.Time `json:"date"`
	Unix int64     `json:"unix"`
}

type PayloadWeather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type StorageType string

const (
	postgresDb StorageType = "postgres"
)

type Storage interface {
	Save(Payload) error
	Get(string) (Payload, error)
	Ping() error
	SaveWeather(PayloadWeather) error
}

func NewStorage(conf conf.Config) (Storage, error) {
	switch StorageType(strings.ToLower(conf.StorageType)) {
	case "postgres":
		storagePostgres, err := NewStoragePostgres(conf.PgAddress)
		if err != nil {
			return nil, err
		}
		return storagePostgres, nil
	}
	return nil, nil
}
