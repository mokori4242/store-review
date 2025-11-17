package repository

import (
	"context"
	"fmt"
	"store-review/internal/domain/store"
	"store-review/internal/infrastructure/gen"
)

type StoreRepository struct {
	queries *sqlc.Queries
}

func NewStoreRepository(queries *sqlc.Queries) store.Repository {
	return &StoreRepository{
		queries: queries,
	}
}

func (r *StoreRepository) FindAll(ctx context.Context) ([]*store.Store, error) {
	result, err := r.queries.GetListStores(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	stores := make([]*store.Store, len(result))
	for i, item := range result {
		stores[i] = &store.Store{
			ID:              item.ID,
			Name:            item.Name,
			RegularHolidays: item.RegularHolidays,
			CategoryNames:   item.CategoryNames,
			PaymentMethods:  item.PaymentMethods,
			WebProfiles:     item.WebProfiles,
		}
	}

	return stores, nil
}
