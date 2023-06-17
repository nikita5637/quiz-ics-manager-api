//go:generate mockery --case underscore --name PlaceStorage --with-expecter

package storage

import (
	"context"

	"github.com/nikita5637/quiz-ics-manager-api/internal/config"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/storage/mysql"
	database "github.com/nikita5637/quiz-ics-manager-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/tx"
)

// PlaceStorage ...
type PlaceStorage interface {
	GetPlaceByExternalPlaceID(ctx context.Context, externalPlaceID int) (*database.Place, error)
}

// NewPlaceStorage ...
func NewPlaceStorage(txManager *tx.Manager) PlaceStorage {
	switch config.GetValue("Driver").String() {
	case config.DriverMySQL:
		return mysql.NewPlaceStorageAdapter(txManager)
	}

	return nil
}
