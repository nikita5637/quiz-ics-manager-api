package model

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/i18n"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ICS files facade errors
var (
	ErrICSFileForGameAlreadyExists = errors.New("ICS file for game already exists")
	ErrICSFileNameAlreadyExists    = errors.New("ICS file name already exists")
	ErrICSFileNotFound             = errors.New("ICS file not found")
)

// GetStatus ...
func GetStatus(ctx context.Context, code codes.Code, err error, reason string, lexeme i18n.Lexeme) *status.Status {
	st := status.New(code, err.Error())
	ei := &errdetails.ErrorInfo{
		Reason: reason,
	}
	lm := &errdetails.LocalizedMessage{
		Locale:  i18n.GetLangFromContext(ctx),
		Message: i18n.GetTranslator(lexeme)(ctx),
	}
	st, err = st.WithDetails(ei, lm)
	if err != nil {
		panic(err)
	}

	return st
}
