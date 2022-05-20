package db

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"timey/context"
	"timey/model"
	"timey/util"
)

const (
	DefaultHost     = "localhost"
	DefaultPort     = "5432"
	DefaultUser     = "postgres"
	DefaultPassword = ""
	DefaultDatabase = "timey"
	DefaultSslMode  = "disable"
	DefaultTimeZone = "Europe/Berlin"
)

func Init() {
	logrus.Info("initialising db")
	dbConfig := New()
	db, err := gorm.Open(postgres.Open(dbConfig.ToConnectionString()))

	if err != nil {
		logrus.Error(err)
		return
	}

	err = db.AutoMigrate(&model.Customer{}, &model.StatementOfWork{}, &model.Activity{})
	if err != nil {
		logrus.Panic(err)
		return
	}

	context.Bind("db", db)
}

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SslMode  string
	TimeZone string
}

func New(options ...ConfigOption) Config {
	config := Config{
		Host:     util.GetEnvWithDefault("DB_HOST", DefaultHost),
		Port:     util.GetEnvWithDefault("DB_PORT", DefaultPort),
		User:     util.GetEnvWithDefault("DB_USER", DefaultUser),
		Password: util.GetEnvWithDefault("DB_PASSWORD", DefaultPassword),
		Database: util.GetEnvWithDefault("DB_NAME", DefaultDatabase),
		SslMode:  util.GetEnvWithDefault("DB_SSL_MODE", DefaultSslMode),
		TimeZone: util.GetEnvWithDefault("DB_TIMEZONE", DefaultTimeZone),
	}

	for _, option := range options {
		option(&config)
	}

	return config
}

func (c Config) ToConnectionString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		c.Host,
		c.User,
		c.Password,
		c.Database,
		c.Port,
		c.SslMode,
		c.TimeZone)
}

type ConfigOption func(*Config)

func WithHost(host string) func(*Config) {
	return func(c *Config) {
		c.Host = host
	}
}

func WithPort(port string) func(*Config) {
	return func(c *Config) {
		c.Port = port
	}
}

func WithUser(user string) func(*Config) {
	return func(c *Config) {
		c.User = user
	}
}

func WithPassword(password string) func(*Config) {
	return func(c *Config) {
		c.Password = password
	}
}

func WithDatabase(database string) func(*Config) {
	return func(c *Config) {
		c.Database = database
	}
}

func WithSslMode(sslMode string) func(*Config) {
	return func(c *Config) {
		c.SslMode = sslMode
	}
}

func WithTimeZone(timeZone string) func(*Config) {
	return func(c *Config) {
		c.TimeZone = timeZone
	}
}
