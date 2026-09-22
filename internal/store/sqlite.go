package store

import (
	"context"

	"duck/internal/api"

	"gorm.io/gorm"
)

// RubberDuck is our database model. It looks a lot like the API model here,
// but that usually doesn't stay true for long :-)
type RubberDuck struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Color string `gorm:"not null"`
	Size  string `gorm:"not null"`
}

// SQLiteStore implements api.DuckStore for a SQLite database using GORM.
// ORMs are a tricky subject in Go. They can be super helpful or extremely painful!
// GORM hides a lot of database details behind some magic, which is fine for this example.
//
// I tend to prefer something like sqlc for simple projects:
// https://github.com/sqlc-dev/sqlc
// Or SQLBoiler when I need more:
// https://github.com/aarondl/sqlboiler
// Both generate Go code from a database schema, much like oapi-codegen does for our API.
type SQLiteStore struct {
	db *gorm.DB
}

// NewSQLiteStore gives us a SQLite store using db.
func NewSQLiteStore(db *gorm.DB) *SQLiteStore {
	return &SQLiteStore{db: db}
}

// Migrate creates or updates the ducks table.
func (s *SQLiteStore) Migrate() error {
	return s.db.AutoMigrate(&RubberDuck{})
}

// Duck gets a single duck.
// NOTE! Duck isn't part of api.DuckStore, but SQLiteStore still implements that
// interface. Interfaces let us add methods here without changing Server.
func (s *SQLiteStore) Duck(ctx context.Context, id uint) (api.RubberDuck, error) {
	duck, err := gorm.G[RubberDuck](s.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return api.RubberDuck{}, err
	}

	return toAPIDuck(duck), nil
}

// ListDucks returns every stored duck.
func (s *SQLiteStore) ListDucks(ctx context.Context) ([]api.RubberDuck, error) {
	stored, err := gorm.G[RubberDuck](s.db).Find(ctx)
	if err != nil {
		return nil, err
	}

	ducks := make([]api.RubberDuck, 0, len(stored))
	for _, duck := range stored {
		ducks = append(ducks, toAPIDuck(duck))
	}
	return ducks, nil
}

// CreateDuck stores and returns a new duck.
func (s *SQLiteStore) CreateDuck(ctx context.Context, duck api.DuckRequest) (api.RubberDuck, error) {
	stored := RubberDuck{
		Name:  duck.Name,
		Color: duck.Color,
		Size:  string(duck.Size),
	}
	if err := gorm.G[RubberDuck](s.db).Create(ctx, &stored); err != nil {
		return api.RubberDuck{}, err
	}

	return toAPIDuck(stored), nil
}

// toAPIDuck is the small bit of copying that keeps our database and API types separate.
func toAPIDuck(duck RubberDuck) api.RubberDuck {
	return api.RubberDuck{
		ID:    int(duck.ID),
		Name:  duck.Name,
		Color: duck.Color,
		Size:  api.RubberDuckSize(duck.Size),
	}
}
