package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/exp/slog"
	"io"
	"log"
	"net/http"
	"os"
	"rabbit/conf"
	db2 "rabbit/db"
	"rabbit/elas"
	"strings"
	"time"
)

func main() {
	// ----- logger part ------
	logger := slog.New(slog.NewJSONHandler(os.Stdout))
	slog.SetDefault(logger)

	// ----- config part ------
	conf := conf.NewConfig()

	// ----- es part ------
	logger.Info("Starting elasticsearch", "es address", conf.EsAddress)
	esClient, err := elas.NewEsClient(conf)
	if err != nil {
		log.Fatal(err)
	}
	es := elas.EsStorage{Client: esClient}

	// ----- db part ------
	logger.Info("Starting db", "pg address", strings.Split(conf.PgAddress, "@")[1])
	db, err := db2.NewStorage(conf)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 1000; i++ {
		// ----- api part -----
		urlWarsaw := fmt.Sprintf(conf.WeatherUrl, conf.Lat["warsaw"], conf.Lon["warsaw"])
		urlTelAviv := fmt.Sprintf(conf.WeatherUrl, conf.Lat["tel_aviv"], conf.Lon["tel_aviv"])
		urlYerevan := fmt.Sprintf(conf.WeatherUrl, conf.Lat["yerevan"], conf.Lon["yerevan"])
		values := []string{urlWarsaw, urlTelAviv, urlYerevan}


		for _, url := range values {
			logger.Info("Sending get request")
			req, _ := http.NewRequest("GET", url, nil)
			res, _ := http.DefaultClient.Do(req)
			go func() {
				defer res.Body.Close()
			}()
			body, err := io.ReadAll(res.Body)
			if err != nil {
				log.Println(err)
			}

			resp := db2.PayloadWeather{}
			json.Unmarshal(body, &resp)

			// ----- save es part ------
			err = es.Save(resp)
			if err != nil {
				log.Println(err)
			}

			//	----- save db part -----
			err = db.SaveWeather(resp)
			if err != nil {
				log.Println(err)
			}
		}
		time.Sleep(160 * time.Second)
	}

}
