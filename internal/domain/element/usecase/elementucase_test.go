package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/billysutomo/chocolate-waffle/internal/domain/mocks"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap"
)

func TestCreateElement(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockElementRepo := mocks.NewMockElementRepository(mockCtrl)

	nowMock := func() time.Time {
		t, _ := time.Parse(time.Kitchen, time.Kitchen)
		return t
	}

	elementModel := domain.ElementModel{
		IDProject: 1,
		Ordernum:  1,
		Type:      "messenger",
		Body:      "body",
		CreatedAt: nowMock(),
		UpdatedAt: nowMock(),
	}

	t.Run("success", func(t *testing.T) {
		mockElementRepo.EXPECT().CreateElement(context.Background(), elementModel).Return(nil)

		a := NewElementUsecase(mockElementRepo, &zap.Logger{})
		isTrue, _ := a.CreateElement(
			context.Background(),
			elementModel.IDProject,
			elementModel.Ordernum,
			elementModel.Type,
			elementModel.Body,
			elementModel.CreatedAt,
			elementModel.UpdatedAt,
		)
		assert.Equal(t, isTrue, true)
	})

}
