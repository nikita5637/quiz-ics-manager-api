package icsfilemanager

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/model"
	icsfilemanagerpb "github.com/nikita5637/quiz-ics-manager-api/pkg/pb/ics_file_manager"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestImplementation_DeleteICSFile(t *testing.T) {
	t.Run("error ICS file not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.icsFilesFacade.EXPECT().DeleteICSFile(fx.ctx, int32(1)).Return(model.ErrICSFileNotFound)

		got, err := fx.implementation.DeleteICSFile(fx.ctx, &icsfilemanagerpb.DeleteICSFileRequest{
			Id: 1,
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("some internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.icsFilesFacade.EXPECT().DeleteICSFile(fx.ctx, int32(1)).Return(errors.New("some error"))

		got, err := fx.implementation.DeleteICSFile(fx.ctx, &icsfilemanagerpb.DeleteICSFileRequest{
			Id: 1,
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.icsFilesFacade.EXPECT().DeleteICSFile(fx.ctx, int32(1)).Return(nil)

		got, err := fx.implementation.DeleteICSFile(fx.ctx, &icsfilemanagerpb.DeleteICSFileRequest{
			Id: 1,
		})

		assert.Equal(t, &emptypb.Empty{}, got)
		assert.NoError(t, err)
	})
}
