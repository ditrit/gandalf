package oauth2

import (
	"io"
	"time"

	"github.com/jinzhu/gorm"
)

// NewConfig create mysql configuration instance
func NewConfig(dsn string, dbType string, tableName string, token bool) *Config {
	return &Config{
		DSN:         dsn,
		DBType:      dbType,
		TableName:   tableName,
		Token:       token,
		MaxLifetime: time.Hour * 2,
	}
}

// Config xorm configuration
type Config struct {
	DSN         string
	DBType      string
	TableName   string
	Token       bool
	MaxLifetime time.Duration
}

// Store mysql token store
type Store struct {
	tableName string
	db        *gorm.DB
	stdout    io.Writer
	ticker    *time.Ticker
}

// NewStore create mysql store instance,
func NewStore(config *Config, gcInterval int) interface{} {
	db, err := gorm.Open(config.DBType, config.DSN)
	if err != nil {
		panic(err)
	}
	if config.Token {
		return NewTokenStoreWithDB(config, db, gcInterval)
	}
	return NewClientStoreWithDB(config, db, gcInterval)
}
