package items

import "errors"

type itemMockRepo struct {
	items map[uint]*Item
}

func NewItemMockRepo() ItemRepo {
	return &itemMockRepo{
		items: make(map[uint]*Item),
	}
}

//Create item
func (repo *itemMockRepo) createItem(item *Item) (*Item, error) {
	if _, ok := repo.items[item.Id]; ok {
		return nil, errors.New("Item with ID already exists")
	}

	repo.items[item.Id] = item

	return item, nil
}

//Read item
func (repo *itemMockRepo) readItem(id uint) (*Item, error) {
	item, ok := repo.items[id]

	if !ok {
		return nil, errors.New("Item with id does not exist")
	}

	return item, nil
}

//Update item
func (repo *itemMockRepo) updateItem(item *Item) (*Item, error) {
	//if item does not exist, create it as an item

	repo.items[item.Id] = item

	return item, nil
}

//Delete item
func (repo *itemMockRepo) deleteItem(id uint) (bool, error) {
	delete(repo.items, id)
	return true, nil
}
