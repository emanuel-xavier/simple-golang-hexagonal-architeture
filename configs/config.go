package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host               string
	Port               string
	User               string
	Pass               string
	Database           string
	MaxIdleConnections int
	MaxOpenConnections int
	SSLMode            string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.pass", "1234")
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.name", "postgres")
	viper.SetDefault("database.max_idle_connection", 5)
	viper.SetDefault("database.max_open_connection", 100)
	viper.SetDefault("database.ssl_mode", "disable")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:               viper.GetString("database.host"),
		Port:               viper.GetString("database.port"),
		User:               viper.GetString("database.user"),
		Pass:               viper.GetString("database.pass"),
		Database:           viper.GetString("database.name"),
		MaxOpenConnections: viper.GetInt("database.max_open_connection"),
		MaxIdleConnections: viper.GetInt("database.max_idle_connection"),
		SSLMode:            viper.GetString("database.ssl_mode"),
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
