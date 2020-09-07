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

// NewClientStoreWithDB
func NewClientStoreWithDB(config *Config, db *gorm.DB, gcInterval int) *ClientStore {
	store := &ClientStore{
		db:        db,
		tableName: "client",
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

	return store
}

// ClientStore
type ClientStore struct {
	tableName string
	db        *gorm.DB
	stdout    io.Writer
	ticker    *time.Ticker
}

// GetByID
func (s *ClientStore) GetByID(context context.Context, id string) (oauth2.ClientInfo, error) {
	if id == "" {
		return nil, nil
	}

	var item models.Client
	if err := s.db.Table(s.tableName).Where("id = ?", id).Find(&item).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}

	return &item, nil
}

// Set
func (s *ClientStore) Set(cli oauth2.ClientInfo) (err error) {
	err = s.db.Table(s.tableName).Save(cli).Error
	var clients models.Client
	s.db.Table(s.tableName).First(&clients)
	fmt.Println("CLIENTS")
	fmt.Println(clients)
	/* for client := range clients {
		fmt.Println(client)
	} */
	return
	//return s.db.Table(s.tableName).Save(cli).Error
}
