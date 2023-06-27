package icsmessage

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/ics"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
	registratorpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
)

// Handle ...
func (h *Handler) Handle(ctx context.Context, event ics.Event) error {
	switch event.Event {
	case ics.EventRegistered:
		_, err := h.icsFilesFacade.GetICSFileByGameID(ctx, event.GameID)
		if err == nil {
			logger.Warnf(ctx, "ICS file for game ID %d already exists", event.GameID)
			return nil
		}

		if !errors.Is(err, model.ErrICSFileNotFound) {
			return fmt.Errorf("get ICS file by game ID error: %w", err)
		}

		game, err := h.registratorServiceClient.GetGameByID(ctx, &registratorpb.GetGameByIDRequest{
			GameId: event.GameID,
		})
		if err != nil {
			return fmt.Errorf("get game by ID error: %w", err)
		}

		pbLeague, err := h.leagueServiceClient.GetLeague(ctx, &leaguepb.GetLeagueRequest{
			Id: game.GetGame().GetLeagueId(),
		})
		if err != nil {
			return fmt.Errorf("get league by ID error: %w", err)
		}

		url := ""
		if pbLeague.GetId() == int32(leaguepb.LeagueID_LEAGUE_ID_QUIZ_PLEASE) {
			url = fmt.Sprintf("https://spb.quizplease.ru/game-page?id=%d", game.GetGame().GetExternalId())
		}

		place, err := h.registratorServiceClient.GetPlaceByID(ctx, &registratorpb.GetPlaceByIDRequest{
			Id: game.GetGame().GetPlaceId(),
		})
		if err != nil {
			return fmt.Errorf("get place by ID error: %w", err)
		}

		address := place.GetPlace().GetAddress()

		placeAddress, err := h.placesFacade.GetAppleAddressByPlaceID(ctx, place.GetPlace().GetId())
		if err != nil {
			logger.Warnf(ctx, "get apple address by place ID error: %s", err.Error())
		} else {
			address = placeAddress
		}

		summary := getSummary(pbLeague.GetName(), game.GetGame().GetNumber())

		generatedICS, err := h.icsGenerator.Generate(summary, address, "", url, game.GetGame().GetDate().AsTime())
		if err != nil {
			return fmt.Errorf("generate ICS error: %w", err)
		}

		// create file begin
		name := uuid.New().String() + h.icsFileExtension
		icsFile, err := os.Create(h.icsFilesFolder + name) //nolint:gosec
		if err != nil {
			return fmt.Errorf("create ICS file error: %w", err)
		}

		_, err = icsFile.Write(generatedICS)
		if err != nil {
			return fmt.Errorf("write data error: %w", err)
		}
		icsFile.Close()
		// create file end

		_, err = h.icsFilesFacade.CreateICSFile(ctx, model.ICSFile{
			GameID: event.GameID,
			Name:   name,
		})
		if err != nil {
			err2 := os.Remove(h.icsFilesFolder + name)
			if err2 != nil {
				return fmt.Errorf("remove file error: %w", err2)
			}

			return fmt.Errorf("create ICS file error: %w", err)
		}

		return nil
	case ics.EventUnregistered:
		icsFile, err := h.icsFilesFacade.GetICSFileByGameID(ctx, event.GameID)
		if err != nil {
			if errors.Is(err, model.ErrICSFileNotFound) {
				logger.Warnf(ctx, "ICS file for game ID %d not found", event.GameID)
				return nil
			}

			return fmt.Errorf("get ICS file by game ID error: %w", err)
		}

		err = h.icsFilesFacade.DeleteICSFile(ctx, icsFile.ID)
		if err != nil {
			return fmt.Errorf("delete ICS file error: %w", err)
		}

		err = os.Remove(h.icsFilesFolder + icsFile.Name)
		if err != nil {
			return fmt.Errorf("remove file error: %w", err)
		}

		return nil
	}

	return fmt.Errorf("invalid event type \"%s\"", event.Event)
}

func getSummary(leagueName, gameNumber string) string {
	return fmt.Sprintf("Игра лиги \"%s\". Номер %s", leagueName, gameNumber)
}
