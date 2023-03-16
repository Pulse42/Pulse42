package store

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
	"github.com/gofiber/storage/sqlite3"
)

// StoreType is the session store, it contains the redis and sqlite3 storage apis
type StoreType struct {
	Sqlite3Storage *sqlite3.Storage
	RedisStorage   *redis.Storage
	Sessions       *session.Store
}

// NewStore creates a new session store based on the storage type
func NewStore(storage string) *StoreType {
	store := &StoreType{}

	switch storage {
	case "sqlite3":
		store.Sqlite3Storage = sqlite3.New()
		store.Sessions = session.New(session.Config{Storage: store.Sqlite3Storage})
	case "redis":
		store.RedisStorage = redis.New()
		store.Sessions = session.New(session.Config{Storage: store.RedisStorage})
	case "memory":
		store.Sessions = session.New()
	default:
		panic("invalid session storage type")
	}

	return store
}

// Store is the session store, it needs to be set before use
var Store *StoreType
