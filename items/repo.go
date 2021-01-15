package items

import (
	"context"
	"database/sql"
	"errors"
	"os"

	"github.com/jackc/pgx/v4"
)

type itemRepo struct {
	db *sql.DB
}

func NewItemRepo(conn string) ItemRepo {
	return &itemRepo{
		db: initDB(conn),
	}
}

func initDB(connStr string) *sql.DB {
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Errorf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	log.Printf("Connected to database: %v\n")
	return conn
}

//Create item
func (repo *itemRepo) createItem(item *Item) (*Item, error) {
	if _, ok := repo.items[item.Id]; ok {
		return nil, errors.New("Item with ID already exists")
	}

	repo.items[item.Id] = item

	return item, nil
}

//Read item
func (repo *itemRepo) readItem(id uint) (*Item, error) {
	item, ok := repo.items[id]

	if !ok {
		return nil, errors.New("Item with id does not exist")
	}

	return item, nil
}

//Update item
func (repo *itemRepo) updateItem(item *Item) (*Item, error) {
	//if item does not exist, create it as an item

	repo.items[item.Id] = item

	return item, nil
}

//Delete item
func (repo *itemRepo) deleteItem(id uint) (bool, error) {
	delete(repo.items, id)
	return true, nil
}
