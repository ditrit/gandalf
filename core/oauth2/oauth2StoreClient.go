package oauth2

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/ditrit/gandalf/core/models"
	oauth2 "github.com/go-oauth2/oauth2/v4"
	"github.com/jinzhu/gorm"
)

var noUpdateContent = "No content found to be updated"

func NewClientStoreWithDB(config *Config, db *gorm.DB, gcInterval int) *ClientStore {
	store := &ClientStore{
		db:        db,
		tableName: "oauth2_client",
		stdout:    os.Stderr,
	}
	if config.TableName != "" {
		store.tableName = config.TableName
	}

	if !db.HasTable(store.tableName) {
		if err := db.Table(store.tableName).CreateTable(&models.Client{}).Error; err != nil {
			panic(err)
		}
	}

	//go store.gc()
	return store
}

// Store mysql token store
type ClientStore struct {
	tableName string
	db        *gorm.DB
	stdout    io.Writer
	ticker    *time.Ticker
}

// GetByID according to the ID for the client information
func (s *ClientStore) GetByID(context context.Context, id string) (oauth2.ClientInfo, error) {
	if id == "" {
		return nil, nil
	}

	var item models.Client
	if err := s.db.Table(s.tableName).Where("id = ?", id).Find(&item).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("BOOP")
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}
	return &item, nil
}

// Set set client information
func (s *ClientStore) Set(id string, cli oauth2.ClientInfo) (err error) {
	return s.db.Table(s.tableName).Save(cli).Error
}
