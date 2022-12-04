package conf

import "github.com/spf13/viper"

type Config struct {
	RabbitAddress string                 `mapstructure:"rabbit"`
	EsAddress     []string               `mapstructure:"es"`
	PgAddress     string                 `mapstructure:"postgres"`
	StorageType   string                 `mapstructure:"storage_type"`
	QueueName     string                 `mapstructure:"queue_name"`
	ApiKey        string                 `mapstructure:"api_key"`
	WeatherUrl    string                 `mapstructure:"weather_url"`
	Lon           map[string]interface{} `mapstructure:"lat"`
	Lat           map[string]interface{} `mapstructure:"lon"`
}

type AddressesType struct {
}

func NewConfig() Config {
	// ----- config part ------
	vi := viper.New()
	vi.SetConfigFile("./conf/app/conf.yaml")
	vi.ReadInConfig()

	return Config{
		RabbitAddress: vi.GetString("RABBIT_ADDRESS"),
		PgAddress:     vi.GetString("POSTGRES_ADDRESS"),
		EsAddress:     vi.GetStringSlice("ES_ADDRESS"),
		StorageType:   vi.GetString("STORAGE_TYPE"),
		QueueName:     vi.GetString("QUEUE_NAME"),
		WeatherUrl:    vi.GetString("WEATHER_URL"),
		Lon:           vi.GetStringMap("LON"),
		Lat:           vi.GetStringMap("LAT"),
	}
}
