package repository

import (
	"belajar-golang-db/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type commentRespositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRespositoryImpl{DB: db}
}

func (repository *commentRespositoryImpl) insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "insert into comments(email,comment) values(?,?)"
	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil

}

func (repository *commentRespositoryImpl) FindbyId(ctx context.Context, id int32) (entity.Comment, error) {
	script := "select id,email,comment from comments where id =? limit 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		//ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("id" + strconv.Itoa(int(id)) + "not found")
	}

}

func (repository *commentRespositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "select id,email,comment from comments"
	rows, err := repository.DB.QueryContext(ctx, script)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
