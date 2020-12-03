package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/billysutomo/chocolate-waffle/internal/domain/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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

		elementUsecase := NewElementUsecase(mockElementRepo, &zap.Logger{})
		isTrue, _ := elementUsecase.CreateElement(
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

func TestGetElementsByIDProject(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockElementRepo := mocks.NewMockElementRepository(mockCtrl)
	nowMock := func() time.Time {
		t, _ := time.Parse(time.Kitchen, time.Kitchen)
		return t
	}
	elementsModel := []domain.ElementModel{
		{
			IDProject: 1,
			Ordernum:  1,
			Type:      "messenger",
			Body:      "",
			CreatedAt: nowMock(),
			UpdatedAt: nowMock(),
		},
		{
			IDProject: 1,
			Ordernum:  2,
			Type:      "messenger",
			Body:      "",
			CreatedAt: nowMock(),
			UpdatedAt: nowMock(),
		},
	}

	t.Run("success", func(t *testing.T) {
		mockElementRepo.EXPECT().GetByIDProject(context.Background(), 1).Return(elementsModel, nil)

		elementUsecase := NewElementUsecase(mockElementRepo, &zap.Logger{})

		elements, err := elementUsecase.GetElementsByIDProject(context.Background(), 1)

		assert.Equal(t, elements, elementsModel)
		assert.NoError(t, err)
	})
}
