package file

import (
	"context"
	"io"
)

type Service interface {
	UploadImage(ctx context.Context, reader io.Reader, size int64, contentType string) (string, error)
	GetImage(ctx context.Context, objectId string) (io.Reader, string, error)
}

type File struct {
	srv Service
}

func NewAPI(srv Service) *File {
	return &File{
		srv: srv,
	}
}
