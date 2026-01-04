package service

import (
	"context"
	"mastery-project/internal/model"
	"mastery-project/internal/repository"
)

type ItemService struct {
	ItemRepo *repository.ItemRepository
}

func NewItemService(itemRepo *repository.ItemRepository) *ItemService {
	return &ItemService{ItemRepo: itemRepo}
}

func (is *ItemService) Save(ctx context.Context, itemReq model.Item) error {
	err := is.ItemRepo.CreateItem(ctx, itemReq)
	if err != nil {
		return err
	}
	return nil
}
func (is *ItemService) GetOne(ctx context.Context, itemId string) (*model.Item, error) {
	item, err := is.ItemRepo.GetItemByID(ctx, itemId)
	if err != nil {
		return nil, err
	}
	return item, nil
}
func (is *ItemService) GetAll(ctx context.Context) ([]model.Item, error) {
	items, err := is.ItemRepo.GetAllItems(ctx)
	if err != nil {
		return nil, err
	}
	return items, nil
}
func (is *ItemService) Delete(ctx context.Context, itemID string) error {
	err := is.ItemRepo.DeleteItemByID(ctx, itemID)
	if err != nil {
		return err
	}
	return nil
}
func (is *ItemService) Update(ctx context.Context, itemId string, item model.UpdateItem) error {
	err := is.ItemRepo.UpdateItemByID(ctx, itemId, item)
	if err != nil {
		return err
	}
	return nil
}
