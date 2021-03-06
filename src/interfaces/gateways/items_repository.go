package gateways

import (
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/domain"
)

// IItemsRepository .
type IItemsRepository interface {
	GetAll(query string) (*domain.Items, error)
}

// ItemsRepository .
type ItemsRepository struct {
	qiitaAPI IQiitaAPI
}

// NewItemsRepository .
func NewItemsRepository(api IQiitaAPI) *ItemsRepository {
	return &ItemsRepository{qiitaAPI: api}
}

// GetAll .
func (r *ItemsRepository) GetAll(query string) (*domain.Items, error) {
	items := &domain.Items{}

	for i := 1; i <= 100; i++ {
		resp, err := r.qiitaAPI.GetItems(i, 100, query)
		if err != nil {
			return nil, err
		}
		if len(*resp) == 0 {
			break
		}

		*items = append(*items, *resp...)
	}

	return items, nil
}
