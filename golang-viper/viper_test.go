package main

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	config := viper.New()
	assert.NotNil(t, config)
}

func TestViperReadConfig(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.json")
	config.ReadInConfig()

	assert.NotNil(t, config)
}

func TestJSON(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")
	err := config.ReadInConfig()
	assert.Nil(t, err)

	name := config.GetString("app.name")
	assert.Equal(t, "Golang Viper", name)

	mysqlUsername := config.GetString("database.mysql.username")
	assert.Equal(t, "localhost", mysqlUsername)

	postgresUsername := config.GetString("database.postgres.username")
	assert.Equal(t, "postgres", postgresUsername)
}

func TestYAML(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.yaml")
	err := config.ReadInConfig()
	assert.Nil(t, err)

	name := config.GetString("app.name")
	assert.Equal(t, "Golang Viper", name)

	mysqlUsername := config.GetString("database.mysql.username")
	assert.Equal(t, "localhost", mysqlUsername)

	postgresUsername := config.GetString("database.postgres.username")
	assert.Equal(t, "postgres", postgresUsername)
}

func TestEnv(t *testing.T) {
	config := viper.New()
	config.SetConfigFile(".env")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	name := config.GetString("APP_NAME")
	assert.Equal(t, "Golang Viper", name)

	mysqlUsername := config.GetString("DATABASE_MYSQL_USERNAME")
	assert.Equal(t, "localhost", mysqlUsername)

	postgresUsername := config.GetString("DATABASE_POSTGRES_USERNAME")
	assert.Equal(t, "postgres", postgresUsername)
}

func TestEnvFile(t *testing.T) {
	config := viper.New()
	config.SetConfigFile(".env")
	config.AutomaticEnv()

	err := config.ReadInConfig()
	assert.Nil(t, err)

	name := config.GetString("APP_NAME")
	assert.Equal(t, "Golang Viper", name)

	mysqlUsername := config.GetString("DATABASE_MYSQL_USERNAME")
	assert.Equal(t, "localhost", mysqlUsername)

	postgresUsername := config.GetString("DATABASE_POSTGRES_USERNAME")
	assert.Equal(t, "postgres", postgresUsername)

	assert.Equal(t, "valen", config.GetString("USERNAME")) // mengambil dari environment variable local komputer
}
