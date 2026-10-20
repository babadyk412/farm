package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

//добавление предмета
func addItem(db *sql.DB, name string, price int, typeId int) {
	query := `INSERT INTO items (name, price, type_id) VALUES (?,?,?)`

	_, err := db.Exec(query, name, price, typeId)
	if err != nil {
		fmt.Println("ошибка при создание предмета", err)
		return
	}
	fmt.Println("создан предмет", name)
}

//получение предмета по ид
func getIdItem(db *sql.DB, id int) (Item, error) {
	query := `SELECT * FROM items where id=?`

	rows, err := db.Query(query, id)
	if err != nil {
		fmt.Println("fff", err)
		return Item{}, err
	}
	var typey Item

	for rows.Next() {

		err := rows.Scan(
			&typey.Id,
			&typey.Name,
			&typey.Type_id,
			&typey.Price,
		)
		if err != nil {
			return typey, err
		}

	}

	if err := rows.Err(); err != nil {
		return Item{}, err
	}
	return typey, nil
}

//получение всех предметов
func getItems(db *sql.DB) ([]Item, error) {
	query := `SELECT * FROM items`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("fff", err)
		return []Item{}, err
	}
	var items []Item

	for rows.Next() {
		var item Item

		err := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Type_id,
			&item.Price,
		)
		if err != nil {
			return []Item{}, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return []Item{}, err
	}
	return items, nil
}