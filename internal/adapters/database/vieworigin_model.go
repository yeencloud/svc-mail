package database

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/yeencloud/bpt-service/internal/domain"
	service "github.com/yeencloud/lib-base"
)

type ViewOriginRepo struct{}

type ViewOrigin struct {
	gorm.Model

	ID string `gorm:"primary_key;unique;not null;default:null;<-:create"`

	IP        string
	Useragent string
}

func (r *ViewOrigin) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New().String()
	return nil
}

func (r ViewOriginRepo) AddViewOrigin(ctx context.Context, origin domain.ViewOrigin) error {
	viewOrigin := toModel(origin)

	return service.WithTransaction(ctx, func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Create(&viewOrigin).Error
	})
}

func (r ViewOriginRepo) CountViewOrigin(ctx context.Context) (int64, error) {
	var count int64

	err := service.WithTransaction(ctx, func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Count(&count).Error
	})

	return count, err
}

func (r ViewOriginRepo) ListViewOrigins(ctx context.Context) ([]domain.ViewOrigin, error) {
	var viewOrigins []ViewOrigin

	err := service.WithTransaction(ctx, func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Find(&viewOrigins).Error
	})

	if err != nil {
		return nil, err
	}

	var origins []domain.ViewOrigin

	for _, viewOrigin := range viewOrigins {
		origins = append(origins, toDomain(viewOrigin))
	}
	return origins, nil
}

func toDomain(viewOrigin ViewOrigin) domain.ViewOrigin {
	return domain.ViewOrigin{
		IP:        viewOrigin.IP,
		Useragent: viewOrigin.Useragent,

		ID: viewOrigin.ID,
	}
}

func toModel(viewOrigin domain.ViewOrigin) ViewOrigin {
	return ViewOrigin{
		IP:        viewOrigin.IP,
		Useragent: viewOrigin.Useragent,

		ID: viewOrigin.ID,
	}
}

func NewViewOriginRepo() ViewOriginRepo {
	return ViewOriginRepo{}
}
