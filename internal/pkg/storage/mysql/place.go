package mysql

import (
	"context"

	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/tx"
)

// PlaceStorageAdapter ...
type PlaceStorageAdapter struct {
	placeStorage *PlaceStorage
}

// NewPlaceStorageAdapter ...
func NewPlaceStorageAdapter(txManager *tx.Manager) *PlaceStorageAdapter {
	return &PlaceStorageAdapter{
		placeStorage: NewPlaceStorage(txManager),
	}
}

// GetPlaceByExternalPlaceID ...
func (a *PlaceStorageAdapter) GetPlaceByExternalPlaceID(ctx context.Context, externalPlaceID int) (*Place, error) {
	return a.placeStorage.GetPlaceByExternalPlaceID(ctx, externalPlaceID)
}
