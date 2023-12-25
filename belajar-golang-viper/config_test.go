package belajar_golang_viper

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	var config = viper.New()
	assert.NotNil(t, config)
}

func TestJson(t *testing.T) {
	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	err := config.ReadInConfig()

	// get value config
	appName := config.GetString("app.name")
	appAuthor := config.GetString("app.author")
	databaseName := config.GetString("database.databaseName")
	fmt.Println(appName)
	fmt.Println(appAuthor)
	fmt.Println(databaseName)
	assert.Nil(t, err)
}

func TestEnvFile(t *testing.T) {
	config := viper.New()

	config.SetConfigFile("config.env")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	err := config.ReadInConfig()

	// get value config
	appName := config.GetString("APP_NAME")
	appAuthor := config.GetString("APP_AUTHOR")
	databaseName := config.GetString("DATABASE_NAME")
	databasePort := config.GetString("DATABASE_PORT")
	fromEnv := config.GetString("FROM_ENV")
	fmt.Println(appName)
	fmt.Println(appAuthor)
	fmt.Println(databaseName)
	fmt.Println(databasePort)
	fmt.Println(fromEnv)
	assert.Nil(t, err)
}
