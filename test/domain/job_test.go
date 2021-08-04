package domain_test

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"project-enconder/domain"
	"testing"
	"time"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo()
	id, _ := uuid.NewV4()
	video.ID = id.String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("path", "Converted", video)
	require.NotNil(t, job)
	require.Nil(t, err)
}
