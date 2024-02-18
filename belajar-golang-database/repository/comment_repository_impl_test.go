package repository

import (
	"context"
	"fmt"
	golang_database "golang-database" //make alias
	"golang-database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConn())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "test@gmail.com",
		Comment: "test comment abc",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	CommentRepository := NewCommentRepository(golang_database.GetConn())

	result, err := CommentRepository.FindById(context.Background(), 24)
	if err != nil {
		panic(err)

	}
	fmt.Println(result)
}

func TestCommentFindAll(t *testing.T) {
	CommentRepository := NewCommentRepository(golang_database.GetConn())

	results, err := CommentRepository.FindAll(context.Background())

	if err != nil {
		panic(err)

	}
	for _, result := range results {
		fmt.Println(result)
	}

}
