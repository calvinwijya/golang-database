package repository

import (
	"belajar-golang-db/entity"
	"context"
)

type CommentRepository interface {
	insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindbyId(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
}
