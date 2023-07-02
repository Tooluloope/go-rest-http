package comment

import (
	"context"
	"fmt"
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
		return Comment{}, err
	}
	return comment, nil
}