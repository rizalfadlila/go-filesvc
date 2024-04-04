package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/rizalfadlila/go-filesvc/model"
	uuid "github.com/satori/go.uuid"
	"time"
)

type (
	ImageRepository interface {
		Save(ctx context.Context, image *model.ImageMetadata) error
	}

	imageModule struct {
		db *sqlx.DB
	}
)

func NewImageModule(db *sqlx.DB) ImageRepository {
	return &imageModule{
		db: db,
	}
}

func (m *imageModule) Save(ctx context.Context, image *model.ImageMetadata) error {
	// use epoc to handle timezone issue
	image.UploadDate = time.Now().UnixMilli()
	image.ID = uuid.NewV4().String()

	const query = `insert into image_metadata (id, filename, filepath, size_file, mimetype, upload_date) values (:id, :filename, :filepath, :size_file, :mimetype, :upload_date)`

	_, err := m.db.NamedExecContext(ctx, query, image)
	return err
}
