package icsfiles

import (
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-ics-manager-api/internal/pkg/storage/mysql"
)

func convertDBICSFileToModelICSFile(dbICSFile database.IcsFile) model.ICSFile {
	return model.ICSFile{
		ID:     int32(dbICSFile.ID),
		GameID: int32(dbICSFile.ExternalGameID),
		Name:   dbICSFile.Name,
	}
}

func convertModelICSFileToDBICSFile(modelICSFile model.ICSFile) database.IcsFile {
	return database.IcsFile{
		ID:             int(modelICSFile.ID),
		ExternalGameID: int(modelICSFile.GameID),
		Name:           modelICSFile.Name,
	}
}
