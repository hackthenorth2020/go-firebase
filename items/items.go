package items

type Item struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

//create interfaces that describe the ItemService and ItemRepo

type ItemService interface {
	CreateItem(*Item) (*Item, error)
	ReadItem(uint) (*Item, error)
	UpdateItem(*Item) (*Item, error)
	DeleteItem(uint) (bool, error)
	ReadAllItems() ([]*Item, error)
}

type ItemRepo interface {
	createItem(*Item) (*Item, error)
	readItem(uint) (*Item, error)
	updateItem(*Item) (*Item, error)
	deleteItem(uint) (bool, error)
	readAllItems() ([]*Item, error)
}
