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

//var noUpdateContent = "No content found to be updated"

func NewStokenItemByTokenInfo(info oauth2.TokenInfo) models.Token {

	item := models.Token{}
	item.SetClientID(info.GetClientID())
	item.SetUserID(info.GetUserID())
	item.SetRedirectURI(info.GetRedirectURI())
	item.SetCode(info.GetCode())
	item.SetCodeCreateAt(info.GetCodeCreateAt())
	item.SetCodeExpiresIn(info.GetCodeExpiresIn())
	item.SetAccess(info.GetAccess())
	item.SetAccessCreateAt(info.GetAccessCreateAt())
	item.SetAccessExpiresIn(info.GetAccessExpiresIn())
	item.SetRefresh(info.GetRefresh())
	item.SetRefreshCreateAt(info.GetRefreshCreateAt())
	item.SetRefreshExpiresIn(info.GetRefreshExpiresIn())

	return item

}

func NewTokenStoreWithDB(config *Config, db *gorm.DB, gcInterval int) *TokenStore {
	store := &TokenStore{
		db:        db,
		tableName: "oauth2_token",
		stdout:    os.Stderr,
	}
	if config.TableName != "" {
		store.tableName = config.TableName
	}
	interval := 600
	if gcInterval > 0 {
		interval = gcInterval
	}
	store.ticker = time.NewTicker(time.Second * time.Duration(interval))

	if !db.HasTable(store.tableName) {
		if err := db.Table(store.tableName).CreateTable(&models.Token{}).Error; err != nil {
			panic(err)
		}
	}

	go store.gc()
	return store
}

// Store mysql token store
type TokenStore struct {
	tableName string
	db        *gorm.DB
	stdout    io.Writer
	ticker    *time.Ticker
}

// Create create and store the new token information
func (s *TokenStore) Create(context context.Context, info oauth2.TokenInfo) error {

	//TODO REVOIR
	item := NewStokenItemByTokenInfo(info)

	if code := info.GetCode(); code != "" {
		item.Code = code
		item.GetCodeCreateAt().Add(info.GetCodeExpiresIn()).Unix()
	} else {
		item.Access = info.GetAccess()
		item.GetAccessCreateAt().Add(info.GetAccessExpiresIn()).Unix()

		if refresh := info.GetRefresh(); refresh != "" {
			item.Refresh = info.GetRefresh()
			item.GetRefreshCreateAt().Add(info.GetRefreshExpiresIn()).Unix()
		}
	}

	return s.db.Table(s.tableName).Create(&item).Error
}

// RemoveByCode delete the authorization code
func (s *TokenStore) RemoveByCode(context context.Context, code string) error {
	return s.db.Table(s.tableName).Where("code = ?", code).Update("code", "").Error
}

// RemoveByAccess use the access token to delete the token information
func (s *TokenStore) RemoveByAccess(context context.Context, access string) error {
	return s.db.Table(s.tableName).Where("access = ?", access).Update("access", "").Error
}

// RemoveByRefresh use the refresh token to delete the token information
func (s *TokenStore) RemoveByRefresh(context context.Context, refresh string) error {
	return s.db.Table(s.tableName).Where("refresh = ?", refresh).Update("refresh", "").Error
}

// GetByCode use the authorization code for token information data
func (s *TokenStore) GetByCode(context context.Context, code string) (oauth2.TokenInfo, error) {
	if code == "" {
		return nil, nil
	}

	var item models.Token
	if err := s.db.Table(s.tableName).Where("code = ?", code).Find(&item).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}

// GetByAccess use the access token for token information data
func (s *TokenStore) GetByAccess(context context.Context, access string) (oauth2.TokenInfo, error) {
	if access == "" {
		return nil, nil
	}

	var item models.Token
	if err := s.db.Table(s.tableName).Where("access = ?", access).Find(&item).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}

// GetByRefresh use the refresh token for token information data
func (s *TokenStore) GetByRefresh(context context.Context, refresh string) (oauth2.TokenInfo, error) {
	if refresh == "" {
		return nil, nil
	}

	var item models.Token
	if err := s.db.Table(s.tableName).Where("refresh = ?", refresh).Find(&item).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}

func (s *TokenStore) gc() {
	for range s.ticker.C {
		now := time.Now().Unix()
		var count int
		if err := s.db.Table(s.tableName).Where("expired_at > ?", now).Count(&count).Error; err != nil {
			s.errorf("[ERROR]:%s\n", err)
			return
		}
		if count > 0 {
			if err := s.db.Table(s.tableName).Where("expired_at > ?", now).Delete(&models.Token{}).Error; err != nil {
				s.errorf("[ERROR]:%s\n", err)
			}
		}
	}
}

// SetStdout set error output
func (s *TokenStore) SetStdout(stdout io.Writer) *TokenStore {
	s.stdout = stdout
	return s
}

// Close close the store
func (s *TokenStore) Close() {
	s.ticker.Stop()
}

func (s *TokenStore) errorf(format string, args ...interface{}) {
	if s.stdout != nil {
		buf := fmt.Sprintf(format, args...)
		s.stdout.Write([]byte(buf))
	}
}
