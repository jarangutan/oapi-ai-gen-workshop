package store

import (
	"context"

	"duck/internal/api"

	"gorm.io/gorm"
)

// RubberDuck here is our database model which looks very similar to our API model
// That's usually not the case :-)
type RubberDuck struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Color string `gorm:"not null"`
	Size  string `gorm:"not null"`
}

// SQLLiteStore implements DuckStore inteface for a sqllite3 database using GORM
// ORMs or Object Relationship Mappers are a tricky subject in Go. They can be super helpful or extremely painful!
// They're particularly notorious for obfuscating a lot of the database away with magic. For our example, that's fine.
//
// I tend to prefer something like [sqlc](https://github.com/sqlc-dev/sqlc) for simple
// or [sqlBoiler](https://github.com/aarondl/sqlboiler) for more complex as these generate code and types from a SQL schema
// similar to how we are doing with oapi-codegen for our API.
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
	ducks, err := gorm.G[RubberDuck](s.db).Find(ctx)
	if err != nil {
		return []api.RubberDuck{}, err
	}

	d := make([]api.RubberDuck, 0, len(ducks))
	for _, duck := range ducks {
		d = append(d, api.RubberDuck{
			Id:    int(duck.ID),
			Name:  duck.Name,
			Color: duck.Color,
			Size:  api.RubberDuckSize(duck.Size),
		})
	}

	return d, nil
}

func (s *SQLLiteStore) CreateDuck(ctx context.Context, duck api.NewRubberDuck) (api.RubberDuck, error) {
	rb := RubberDuck{Name: duck.Name, Color: duck.Color, Size: string(duck.Size)}
	result := gorm.WithResult()
	err := gorm.G[RubberDuck](s.db, result).Create(ctx, &rb)
	if err != nil {
		return api.RubberDuck{}, err
	}

	return api.RubberDuck{
		Id:    int(rb.ID),
		Color: rb.Color,
		Name:  rb.Name,
		Size:  api.RubberDuckSize(rb.Size),
	}, nil
}
