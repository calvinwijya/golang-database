package repository

import (
	belajargolangdb "belajar-golang-db"
	"belajar-golang-db/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	CommentRepository := NewCommentRepository(belajargolangdb.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "test repository",
	}
	result, err := CommentRepository.insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	CommentRepository := NewCommentRepository(belajargolangdb.GetConnection())
	ctx := context.Background()
	comment, err := CommentRepository.FindbyId(ctx, 13)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	CommentRepository := NewCommentRepository(belajargolangdb.GetConnection())
	ctx := context.Background()
	comments, err := CommentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
