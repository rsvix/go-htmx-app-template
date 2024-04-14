package cookiestore

import (
	"time"

	"github.com/gorilla/securecookie"
	"github.com/wader/gormstore/v2"
	"gorm.io/gorm"
)

// Store session cookies
func Start(db *gorm.DB) *gormstore.Store {
	store := gormstore.NewOptions(
		db, // *gorm.DB
		gormstore.Options{
			TableName:       "sessions", // "sessions" is default
			SkipCreateTable: false,      // false is default
		},
		[]byte(securecookie.GenerateRandomKey(64)), // 32 or 64 bytes recommended, required
		[]byte(securecookie.GenerateRandomKey(32)), // nil, 16, 24 or 32 bytes, optional
	)

	store.SessionOpts.Path = "/"
	store.SessionOpts.MaxAge = 86400
	store.SessionOpts.HttpOnly = true
	store.SessionOpts.Secure = false

	quit := make(chan struct{})
	go store.PeriodicCleanup(1*time.Hour, quit)

	return store
}
