package store

import (
	"context"

	"duck/internal/api"

	"gorm.io/gorm"
)

type RubberDuck struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Color string `gorm:"not null"`
	Size  string `gorm:"not null"`
}

type SQLLiteStore struct {
	db *gorm.DB
}

func NewSQLiteStore(db *gorm.DB) *SQLLiteStore {
	return &SQLLiteStore{
		db: db,
	}
}

func (s *SQLLiteStore) Migrate() {
	// Migrate the schema
	s.db.AutoMigrate(&RubberDuck{})
}

func (s *SQLLiteStore) GetDucks(ctx context.Context) ([]api.RubberDuck, error) {
	return nil, nil
}

func (s *SQLLiteStore) CreateDuck(ctx context.Context, duck api.NewRubberDuck) (api.RubberDuck, error) {
	rb := RubberDuck{Name: duck.Name, Color: duck.Color, Size: string(duck.Size)}
	result := gorm.WithResult()
	err := gorm.G[RubberDuck](s.db, result).Create(ctx, &rb)
	if err != nil {
		return api.RubberDuck{}, nil
	}

	return api.RubberDuck{
		Id:    int(rb.ID),
		Color: rb.Color,
		Name:  rb.Name,
		Size:  api.RubberDuckSize(rb.Size),
	}, nil
}
