package config

import (
	"fmt"
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

var cfg Config

type Config struct {
	EmailConfig struct {
		SmtpRelayAddress  string `env:"SMTP_RELAY_ADDRESS,default=smtp-relay.brevo.com"`
		SmtpPort          int    `env:"SMTP_PORT,default=587"`
		FromEmailAddress  string `env:"FROM_EMAIL_ADDRESS,default=muhammad.danish1492@gmail.com"`
		FromEmailPassword string `env:"FROM_EMAIL_PASSWORD,default=KJVqdw8TS5zfZIOt"`
		//SmsApiKey         string `env:"SMS_API_KEY,default=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"`
		SmsApiKey string `env:"SMS_API_KEY,default=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"`
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load environment variables: " + err.Error())
		fmt.Println("Using default values")
	}

	err = envdecode.Decode(&cfg)
	if err != nil {
		fmt.Println("Failed to parse configurations: " + err.Error())
	}
}

func GetAppConfig() Config {
	return cfg
}
