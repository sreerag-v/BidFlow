package config

import (
	"log"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost                string `mapstructure:"DB_HOST"`
	DBName                string `mapstructure:"DB_NAME"`
	DBUser                string `mapstructure:"DB_USER"`
	DBPort                string `mapstructure:"DB_PORT"`
	DBPassword            string `mapstructure:"DB_PASSWORD"`
	AWS_REGION            string `mapstructure:"AWS_REGION"`
	AWS_ACCESS_KEY_ID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWS_SECRET_ACCESS_KEY string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
}

type Twilio struct {
	AuthToken  string `mapstructure:"AuthToken"`
	AccountSid string `mapstructure:"AccountSid"`
	ServiceSid string `mapstructure:"ServiceSid"`
}
type Smtp struct {
	Email    string `mapstructure:"Email"`
	Password string `mapstructure:"Password"`
}
var envs = []string{
	"DB_HOST",
	"DB_NAME",
	"DB_USER",
	"DB_PORT",
	"DB_PASSWORD",
	"AWS_REGION",
	"AWS_ACCESS_KEY_ID",
	"AWS_SECRET_ACCESS_KEY",

	"authtoken", "accountsid", "servicesid", // twilio
	"Email", "Password", //smtp

}
var twilio Twilio
var smtp Smtp 

func LoadConfig()(Config,error){
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	err:=viper.ReadInConfig()

	if err!=nil{
		log.Fatal("error while loading viper")
	}

	for _,env := range envs{
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&twilio); err != nil {
		return config, err
	}
	if err := viper.Unmarshal(&smtp); err != nil {
		return config, err
	}

	return config, nil
}

func GetTilio() Twilio {
	return twilio
}
func GetSmtp() Smtp {
	return smtp
}