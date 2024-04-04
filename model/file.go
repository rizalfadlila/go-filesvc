package model

type FileUpload struct {
}

type ImageMetadata struct {
	ID         string `db:"id"`
	Filename   string `db:"filename"`
	Filepath   string `db:"filepath"`
	Size       int64  `db:"size_file"`
	Mimetype   string `db:"mimetype"`
	UploadDate int64  `db:"upload_date"`
}

func (m *ImageMetadata) TableName() string {
	return "image_metadata"
}
