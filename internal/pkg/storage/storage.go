package storage

import (
	"context"
	"database/sql"
	"errors"

	"github.com/nikita5637/quiz-ics-manager-api/internal/config"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/storage/mysql"
)

// NewDB ...
func NewDB(ctx context.Context, driver string) (*sql.DB, error) {
	switch driver {
	case config.DriverMySQL:
		return mysql.NewDB(ctx, config.GetMySQLDatabaseDSN())
	}

	return nil, errors.New("unknown driver")
}
