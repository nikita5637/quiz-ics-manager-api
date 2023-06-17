package mysql

import (
	"context"

	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/tx"
)

// ICSFileStorageAdapter ...
type ICSFileStorageAdapter struct {
	icsFileStorage *IcsFileStorage
}

// NewICSFileStorageAdapter ...
func NewICSFileStorageAdapter(txManager *tx.Manager) *ICSFileStorageAdapter {
	return &ICSFileStorageAdapter{
		icsFileStorage: NewIcsFileStorage(txManager),
	}
}

// CreateICSFile ...
func (a *ICSFileStorageAdapter) CreateICSFile(ctx context.Context, dbICSFile IcsFile) (int, error) {
	id, err := a.icsFileStorage.Insert(ctx, dbICSFile)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// DeleteICSFile ...
func (a *ICSFileStorageAdapter) DeleteICSFile(ctx context.Context, id int) error {
	return a.icsFileStorage.Delete(ctx, id)
}

// GetICSFileByID ...
func (a *ICSFileStorageAdapter) GetICSFileByID(ctx context.Context, id int) (*IcsFile, error) {
	return a.icsFileStorage.GetIcsFileByID(ctx, id)
}

// GetICSFileByGameID ...
func (a *ICSFileStorageAdapter) GetICSFileByExternalGameID(ctx context.Context, externalGameID int) (*IcsFile, error) {
	return a.icsFileStorage.GetIcsFileByExternalGameID(ctx, externalGameID)
}

// GetICSFiles ...
func (a *ICSFileStorageAdapter) GetICSFiles(ctx context.Context) ([]IcsFile, error) {
	return a.icsFileStorage.Find(ctx, nil, "id")
}
