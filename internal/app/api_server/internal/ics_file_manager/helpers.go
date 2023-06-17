package icsfilemanager

import (
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/model"
	icsfilemanagerpb "github.com/nikita5637/quiz-ics-manager-api/pkg/pb/ics_file_manager"
)

var (
	icsFileNotFoundLexeme = i18n.Lexeme{
		Key:      "ICS_file_not_found",
		FallBack: "ICS file not found",
	}
)

func convertModelICSFileToProtoICSFile(modelICSFile model.ICSFile) *icsfilemanagerpb.ICSFile {
	return &icsfilemanagerpb.ICSFile{
		Id:     modelICSFile.ID,
		GameId: modelICSFile.GameID,
		Name:   modelICSFile.Name,
	}
}
