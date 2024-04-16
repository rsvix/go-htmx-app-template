package cookiestore

import (
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/wader/gormstore/v2"
	"gorm.io/gorm"
)

func Start(db *gorm.DB) *gormstore.Store {
	store := gormstore.NewOptions(
		db,
		gormstore.Options{
			TableName:       "sessions",
			SkipCreateTable: false,
		},
		[]byte(securecookie.GenerateRandomKey(64)), // 32 or 64 bytes recommended, required
		[]byte(securecookie.GenerateRandomKey(32)), // nil, 16, 24 or 32 bytes, optional
	)

	store.SessionOpts.Path = "/"
	store.SessionOpts.MaxAge = 86400
	store.SessionOpts.HttpOnly = true
	store.SessionOpts.Secure = false
	store.SessionOpts.SameSite = http.SameSiteLaxMode
	// store.SessionOpts.Domain = ""

	quit := make(chan struct{})
	go store.PeriodicCleanup(1*time.Hour, quit)

	return store
}
