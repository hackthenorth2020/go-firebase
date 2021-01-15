package items

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

type itemRepo struct {
	// db *pgx.Conn
	connStr string
}

func NewItemRepo(conn string) ItemRepo {
	return &itemRepo{
		// db: initDB(conn),
		connStr: conn,
	}
}

func initDB(connStr string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	defer log.Printf("Conn closing")
	log.Printf("Connected to database\n")
	return conn
}

//Create item
func (repo *itemRepo) createItem(item *Item) (*Item, error) {
	// err := repo.db.QueryRow(context.Background(), "INSERT INTO items (name, owner) VALUES ($1, $2) RETURNING id", &item.Name, &item.Owner).Scan(&item.Id)
	conn, err := pgx.Connect(context.Background(), repo.connStr)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	err = conn.QueryRow(context.Background(), "INSERT INTO items (name, owner) VALUES ($1, $2) RETURNING id", &item.Name, &item.Owner).Scan(&item.Id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

//Read item
func (repo *itemRepo) readItem(id uint) (*Item, error) {
	conn, err := pgx.Connect(context.Background(), repo.connStr)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	item := &Item{}
	row := conn.QueryRow(context.Background(), "SELECT * FROM items WHERE id = $1", id)
	err = row.Scan(&item.Id, &item.Name, &item.Owner)
	if err != nil {
		return nil, err
	}

	return item, nil
}

//Update item
func (repo *itemRepo) updateItem(item *Item) (*Item, error) {
	conn, err := pgx.Connect(context.Background(), repo.connStr)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "UPDATE (name, owner) SET name = $2, owner = $3 FROM items WHERE id = $1", &item.Id, &item.Name, &item.Owner)
	if err != nil {
		return nil, err
	}

	return item, nil
}

//Delete item
func (repo *itemRepo) deleteItem(id uint) (bool, error) {
	conn, err := pgx.Connect(context.Background(), repo.connStr)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "DELETE FROM items WHERE id = $1", id)
	if err != nil {
		return false, err
	}

	return true, nil
}
