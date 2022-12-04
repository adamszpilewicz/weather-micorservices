package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"golang.org/x/exp/slog"
	"os"
)

type StoragePostgres struct {
	conn   *pgx.Conn
	logger  *slog.Logger
	address string
}

func NewStoragePostgres(address string) (StoragePostgres, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout))
	conn, err := pgx.Connect(context.Background(), address)
	if err != nil {
		return StoragePostgres{}, err
	}
	return StoragePostgres{conn: conn, logger: logger, address: address}, nil

}

func (sp StoragePostgres) Ping() error {
	sp.logger.Info("Pinging database", "db url", sp.address)
	return sp.conn.Ping(context.Background())
}

func (sp StoragePostgres) Save(payload Payload) error {
	sp.logger.Info("Saving into database", "payload", payload)
	_, err := sp.conn.Exec(
		context.Background(),
		"insert into users(name, age, date, unix) values($1, $2, $3, $4)",
		payload.Name, payload.Age, payload.Date, payload.Unix,
	)
	return err
}

func (sp StoragePostgres) SaveWeather(payload PayloadWeather) error {
	sp.logger.Info("Saving weather into database", "payload", payload)
	_, err := sp.conn.Exec(
		context.Background(),
		"insert into weather(unix_date, temperature, temperature_feel, unix_sunrise, unix_sunset, sky, city) values($1, $2, $3, $4, $5, $6, $7)",
		payload.Dt, payload.Main.Temp, payload.Main.FeelsLike, payload.Sys.Sunrise, payload.Sys.Sunset, payload.Weather[0].Main, payload.Name,
	)
	return err
}

func (sp StoragePostgres) Get(name string) (Payload, error) {
	return Payload{}, nil
}
