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

func TestGet(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProjectRepo := mocks.NewMockProjectRepository(mockCtrl)
	nowMock := func() time.Time {
		t, _ := time.Parse(time.Kitchen, time.Kitchen)
		return t
	}
	projectsModel := []domain.ProjectModel{
		{
			ID:             1,
			IDUser:         13,
			URL:            "url",
			ProfilePicture: "profile",
			Title:          "title",
			Description:    "desc",
			CreatedAt:      nowMock(),
			UpdatedAt:      nowMock(),
		},
		{
			ID:             2,
			IDUser:         13,
			URL:            "url",
			ProfilePicture: "profile",
			Title:          "title",
			Description:    "desc",
			CreatedAt:      nowMock(),
			UpdatedAt:      nowMock(),
		},
	}

	t.Run("success", func(t *testing.T) {
		mockProjectRepo.EXPECT().Get(context.Background(), 1).Return(projectsModel, nil)

		projectUsecase := NewProjectUsecase(mockProjectRepo, &zap.Logger{})

		projects, err := projectUsecase.Get(context.Background(), 1)

		assert.Equal(t, projects, projectsModel)
		assert.NoError(t, err)
	})
}
