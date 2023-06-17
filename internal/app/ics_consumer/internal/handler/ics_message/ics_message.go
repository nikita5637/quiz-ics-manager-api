//go:generate mockery --case underscore --name ICSFilesFacade --with-expecter
//go:generate mockery --case underscore --name ICSGenerator --with-expecter
//go:generate mockery --case underscore --name PlacesFacade --with-expecter
//go:generate mockery --case underscore --name RegistratorServiceClient --with-expecter

package icsmessage

import (
	"context"
	"time"

	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/model"
	registratorpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
)

// ICSFilesFacade ...
type ICSFilesFacade interface {
	CreateICSFile(ctx context.Context, icsFile model.ICSFile) (model.ICSFile, error)
	DeleteICSFile(ctx context.Context, id int32) error
	GetICSFileByGameID(ctx context.Context, gameID int32) (model.ICSFile, error)
}

// ICSGenerator ...
type ICSGenerator interface {
	Generate(summary, address, description, url string, gameDateTime time.Time) ([]byte, error)
}

// PlacesFacade ...
type PlacesFacade interface {
	GetAppleAddressByPlaceID(ctx context.Context, placeID int32) (string, error)
}

// RegistratorServiceClient ...
type RegistratorServiceClient interface {
	registratorpb.RegistratorServiceClient
}

// Handler ...
type Handler struct {
	icsFileExtension         string
	icsFilesFacade           ICSFilesFacade
	icsFilesFolder           string
	icsGenerator             ICSGenerator
	placesFacade             PlacesFacade
	registratorServiceClient RegistratorServiceClient
}

// Config ...
type Config struct {
	IcsFileExtension         string
	ICSFilesFacade           ICSFilesFacade
	IcsFilesFolder           string
	ICSGenerator             ICSGenerator
	PlacesFacade             PlacesFacade
	RegistratorServiceClient RegistratorServiceClient
}

// New ...
func New(cfg Config) *Handler {
	return &Handler{
		icsFileExtension:         cfg.IcsFileExtension,
		icsFilesFacade:           cfg.ICSFilesFacade,
		icsFilesFolder:           cfg.IcsFilesFolder,
		icsGenerator:             cfg.ICSGenerator,
		placesFacade:             cfg.PlacesFacade,
		registratorServiceClient: cfg.RegistratorServiceClient,
	}
}
