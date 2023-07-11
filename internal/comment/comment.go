package comment

import (
	"context"
	"errors"
)

var (
	ErrFetchingComment = errors.New("error fetching comment")
	ErrNotImplemented  = errors.New("not implemented")
)

type Store interface {
	GetComment(ctx context.Context, id string) (Comment, error)
	PostComment(ctx context.Context, comment Comment) (Comment, error)
	DeleteComment(ctx context.Context, id string) error
	UpdateComment(ctx context.Context, comment Comment) (Comment, error)
}

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	comment, err := s.store.GetComment(ctx, id)
	if err != nil {
		return Comment{}, ErrFetchingComment
	}
	return comment, nil
}

func (s *Service) UpdateComment(ctx context.Context, comment Comment) (Comment, error) {

	comm, err := s.store.UpdateComment(ctx, comment)

	if err != nil {
		return Comment{}, err
	}
	return comm, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {

	return s.store.DeleteComment(ctx, id)

}

func (s *Service) PostComment(ctx context.Context, comment Comment) (Comment, error) {

	insertedComment, err := s.store.PostComment(ctx, comment)

	if err != nil {
		return Comment{}, err
	}

	return insertedComment, nil
}
