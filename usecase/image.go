package usecase

import (
	"context"
	"github.com/rizalfadlila/go-filesvc/model"
	"github.com/rizalfadlila/go-filesvc/repository"
)

type (
	imageModule struct {
		imageRepo repository.ImageRepository
	}

	ImageUsecase interface {
		Save(ctx context.Context, image *model.ImageMetadata) error
	}
)

func NewImageUsecase(imageRepo repository.ImageRepository) ImageUsecase {
	return &imageModule{
		imageRepo: imageRepo,
	}
}

func (u *imageModule) Save(ctx context.Context, image *model.ImageMetadata) error {
	return u.imageRepo.Save(ctx, image)
}
