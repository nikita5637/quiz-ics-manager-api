package places

import (
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/tx"
)

// Facade ...
type Facade struct {
	db           *tx.Manager
	placeStorage storage.PlaceStorage
}

// Config ...
type Config struct {
	PlaceStorage storage.PlaceStorage
	TxManager    *tx.Manager
}

// New ...
func New(cfg Config) *Facade {
	return &Facade{
		db:           cfg.TxManager,
		placeStorage: cfg.PlaceStorage,
	}
}
