package cookiestore

import (
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/wader/gormstore/v2"
	"gorm.io/gorm"
)

// https://www.alexedwards.net/blog/working-with-cookies-in-go

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
	store.SessionOpts.Domain = "localhost" //"example.com"
	store.SessionOpts.MaxAge = 86400
	store.SessionOpts.Secure = false
	store.SessionOpts.HttpOnly = true
	store.SessionOpts.SameSite = http.SameSiteStrictMode

	quit := make(chan struct{})
	go store.PeriodicCleanup(1*time.Hour, quit)

	return store
}
