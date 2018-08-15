package service

import (
	"context"
	"errors"
)

type StringService struct {
	// here i can cache db connections and similar stuff
}

func New() (*StringService, error) {
	return &StringService{}, nil
}

func (svc *StringService) Count(ctx context.Context, input string) (int, error) {
	if len(input) == 0 {
		return 0, errors.New("string can't be empty")
	}
	return len(input), nil
}
