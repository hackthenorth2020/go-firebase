package items

import "log"

type itemService struct {
	repo ItemRepo
}

func NewItemService(conn string) ItemService {
	return &itemService{
		repo: NewItemRepo(conn),
	}
}

func (srv *itemService) CreateItem(item *Item) (*Item, error) {
	log.Printf("[ITEM SRV] [CREATE ITEM] %v", item)
	result, err := srv.repo.createItem(item)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (srv *itemService) ReadItem(id uint) (*Item, error) {
	log.Printf("[ITEM SRV] [READ ITEM] %v", id)
	result, err := srv.repo.readItem(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (srv *itemService) UpdateItem(item *Item) (*Item, error) {
	log.Printf("[ITEM SRV] [UPDATE ITEM] %v", item)
	result, err := srv.repo.updateItem(item)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (srv *itemService) DeleteItem(id uint) (bool, error) {
	log.Printf("[ITEM SRV] [DELETE ITEM] %v", id)
	result, err := srv.repo.deleteItem(id)
	if err != nil {
		return false, err
	}
	return result, nil
}
