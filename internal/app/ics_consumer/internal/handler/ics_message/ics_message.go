//go:generate mockery --case underscore --name ICSFilesFacade --with-expecter
//go:generate mockery --case underscore --name ICSGenerator --with-expecter
//go:generate mockery --case underscore --name PlacesFacade --with-expecter
//go:generate mockery --case underscore --name LeagueServiceClient --with-expecter
//go:generate mockery --case underscore --name PlaceServiceClient --with-expecter
//go:generate mockery --case underscore --name GameServiceClient --with-expecter

package icsmessage

import (
	"context"
	"time"

	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/model"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
	placepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/place"
	"google.golang.org/grpc"
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

// GameServiceClient ...
type GameServiceClient interface {
	GetGame(ctx context.Context, in *gamepb.GetGameRequest, opts ...grpc.CallOption) (*gamepb.Game, error)
}

// LeagueServiceClient ...
type LeagueServiceClient interface {
	GetLeague(ctx context.Context, in *leaguepb.GetLeagueRequest, opts ...grpc.CallOption) (*leaguepb.League, error)
}

// PlaceServiceClient ...
type PlaceServiceClient interface {
	GetPlace(ctx context.Context, in *placepb.GetPlaceRequest, opts ...grpc.CallOption) (*placepb.Place, error)
}

// Handler ...
type Handler struct {
	icsFileExtension string
	icsFilesFacade   ICSFilesFacade
	icsFilesFolder   string
	icsGenerator     ICSGenerator
	placesFacade     PlacesFacade

	gameServiceClient   GameServiceClient
	leagueServiceClient LeagueServiceClient
	placeServiceClient  PlaceServiceClient
}

// Config ...
type Config struct {
	IcsFileExtension string
	ICSFilesFacade   ICSFilesFacade
	IcsFilesFolder   string
	ICSGenerator     ICSGenerator
	PlacesFacade     PlacesFacade

	GameServiceClient   GameServiceClient
	LeagueServiceClient LeagueServiceClient
	PlaceServiceClient  PlaceServiceClient
}

// New ...
func New(cfg Config) *Handler {
	return &Handler{
		icsFileExtension: cfg.IcsFileExtension,
		icsFilesFacade:   cfg.ICSFilesFacade,
		icsFilesFolder:   cfg.IcsFilesFolder,
		icsGenerator:     cfg.ICSGenerator,
		placesFacade:     cfg.PlacesFacade,

		gameServiceClient:   cfg.GameServiceClient,
		leagueServiceClient: cfg.LeagueServiceClient,
		placeServiceClient:  cfg.PlaceServiceClient,
	}
}
