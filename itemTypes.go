package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

//добавление типа
func addItemType(db *sql.DB, name string) {
	query := `INSERT INTO items_type (name) VALUES (?)`

	fmt.Println("создан тип", name)

	_, err := db.Exec(query, name)
	if err != nil {
		fmt.Println("ошибка при создание типа", err)
		return
	}
}

//получение типа по ид
func getTypeId(db *sql.DB, id int) (Type_id, error) {
	query := `SELECT * FROM items where id=?`

	rows, err := db.Query(query, id)
	if err != nil {
		fmt.Println("fff", err)
		return Type_id{}, err
	}
	var typey Type_id

	for rows.Next() {

		err := rows.Scan(
			&typey.Name,
		)
		if err != nil {
			return Type_id{}, err
		}

	}

	if err := rows.Err(); err != nil {
		return Type_id{}, err
	}
	return typey, nil
}

//получение всех типов
func getTypes(db *sql.DB) ([]Type_id, error) {
	query := `SELECT * FROM items`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("fff", err)
		return []Type_id{}, err
	}
	var types []Type_id

	for rows.Next() {
		var typey Type_id

		err := rows.Scan(
			&typey.Name,
		)
		if err != nil {
			return []Type_id{}, err
		}

		types = append(types, typey)
	}

	if err := rows.Err(); err != nil {
		return []Type_id{}, err
	}
	return types, nil
}