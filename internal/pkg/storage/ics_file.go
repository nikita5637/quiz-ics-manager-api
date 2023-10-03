//go:generate mockery --case underscore --name ICSFileStorage --with-expecter

package storage

import (
	"context"

	"github.com/nikita5637/quiz-ics-manager-api/internal/config"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/storage/mysql"
	database "github.com/nikita5637/quiz-ics-manager-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/tx"
)

// ICSFileStorage ...
type ICSFileStorage interface {
	CreateICSFile(ctx context.Context, dbICSFile database.IcsFile) (int, error)
	DeleteICSFile(ctx context.Context, id int) error
	GetICSFileByID(ctx context.Context, id int) (*database.IcsFile, error)
	GetICSFileByExternalGameID(ctx context.Context, externalGameID int) (*database.IcsFile, error)
	GetICSFiles(ctx context.Context) ([]database.IcsFile, error)
}

// NewICSFileStorage ...
func NewICSFileStorage(driver string, txManager *tx.Manager) ICSFileStorage {
	switch driver {
	case config.DriverMySQL:
		return mysql.NewICSFileStorageAdapter(txManager)
	}

	return nil
}
