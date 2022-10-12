package repository

import (
	"context"
	"fmt"
	"reviuw"
	"reviuw/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsert(t *testing.T) {
	CommentRepository := NewCommentRepository(reviuw.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Name:    "Arif",
		Email:   "arif@gmail.com",
		Comment: "Comment Aja",
	}

	for i := 0; i < 10; i++ {
		result, err := CommentRepository.Insert(ctx, comment)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	}
	fmt.Println("Success Insert New Comment")
}

func TestFindById(t *testing.T) {
	CommentRepository := NewCommentRepository(reviuw.GetConnection())

	comment, err := CommentRepository.FindById(context.Background(), 20)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	CommentRepository := NewCommentRepository(reviuw.GetConnection())

	comments, err := CommentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
