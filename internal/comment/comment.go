package comment

import (
	"context"
	"errors"
	"fmt"
)
var (
	ErrFetchingComment = errors.New("error fetching comment")
	ErrNotImplemented = errors.New("not implemented")
)

type Store interface {
	GetComment(ctx context.Context, id string) (Comment, error)
}

type Comment struct {
	ID string
	Slug string
	Body string
	Author string
}

type Service struct{
	store Store
}


func NewService(store Store ) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("GetComment")
	comment, err := s.store.GetComment(ctx, id)
	if err != nil {
		return Comment{}, ErrFetchingComment
	}
	return comment, nil
}

func (s *Service) UpdateComment(ctx context.Context, comment Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, comment Comment) (Comment, error) {
	return Comment{} , ErrNotImplemented
}