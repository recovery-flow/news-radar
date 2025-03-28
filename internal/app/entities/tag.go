package entities

import (
	"context"

	"github.com/recovery-flow/news-radar/internal/app/models"
	"github.com/recovery-flow/news-radar/internal/config"
	"github.com/recovery-flow/news-radar/internal/data"
)

type tagsRepo interface {
	Create(ctx context.Context, tag models.Tag) error
	Delete(ctx context.Context, name string) error
	Update(ctx context.Context, name string, fields map[string]any) error
	Get(ctx context.Context, name string) (*models.Tag, error)
}

type Tags struct {
	data tagsRepo
}

func NewTags(cfg config.Config) (*Tags, error) {
	repo, err := data.NewTags(cfg)
	if err != nil {
		return nil, err
	}

	return &Tags{
		data: repo,
	}, nil
}
