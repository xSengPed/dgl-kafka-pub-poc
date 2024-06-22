package config

type Kafka struct {
	HostName string
}

type Config struct {
	Kafka *Kafka
}

func NewConfig() *Config {
	config := new(Config)

	return config
}
