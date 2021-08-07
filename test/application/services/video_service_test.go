package services_test

import (
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"log"
	"project-enconder/application/repositories"
	"project-enconder/application/services"
	"project-enconder/domain"
	"project-enconder/infrastructure/database"
	"testing"
	"time"
)

func init() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func prepare() (*domain.Video, repositories.VideoRepositoryDb) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	id, _ := uuid.NewV4()
	video.ID = id.String()
	video.FilePath = "filhodomato.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}

	return video, repo
}

func TestVideoServiceDownload(t *testing.T) {
	video, repo := prepare()

	videoService := services.NewVideoService()
	videoService.Video = video
	videoService.VideoRepository = repo

	err := videoService.Download("encondertest")
	require.Nil(t, err)

	err = videoService.Fragment()
	require.Nil(t, err)
}
