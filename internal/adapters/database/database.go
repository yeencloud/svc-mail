package database

import (
	"context"

	"gorm.io/gorm"
)

type Database struct {
	Gorm *gorm.DB
}

func NewDatabase(ctx context.Context, gorm *gorm.DB) (*Database, error) {
	db := &Database{
		Gorm: gorm,
	}

	if err := db.registerModels(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func (r Database) registerModels(ctx context.Context) error {
	return r.Gorm.WithContext(ctx).AutoMigrate(&ViewOrigin{})
}
