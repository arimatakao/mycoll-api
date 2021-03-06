package apiserver

type Config struct {
	BindAddr  string `toml:"bind_addr"`
	LogLevel  string `toml:"log_lvl"`
	DBURI     string `toml:"db_uri"`
	SecretKey string `toml:"secret_key"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:  ":8000",
		LogLevel:  "debug",
		DBURI:     "",
		SecretKey: "1111",
	}
}
