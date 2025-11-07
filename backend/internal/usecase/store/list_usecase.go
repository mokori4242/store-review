package store

import (
	"context"
	"errors"
	"store-review/internal/domain/store"
)

type ListOutput struct {
	Stores []*store.Store
}

type ListUseCase struct {
	storeRepo store.Repository
}

func NewListUseCase(storeRepo store.Repository) *ListUseCase {
	return &ListUseCase{
		storeRepo: storeRepo,
	}
}

func (lc *ListUseCase) Execute(ctx context.Context) (*ListOutput, error) {
	stores, err := lc.storeRepo.FindAll(ctx)
	if err != nil {
		return nil, errors.New("店舗を取得できませんでした。")
	}

	return &ListOutput{
		Stores: stores,
	}, nil
}
