package repositories_test

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"project-enconder/application/repositories"
	"project-enconder/domain"
	"project-enconder/infrastructure/database"
	"testing"
	"time"
)

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	id, _ := uuid.NewV4()
	video.ID = id.String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID, video.ID)
}