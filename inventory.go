package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)
//добавление предмета в инвентарь
func addItemInventory(db *sql.DB, idUser int, idItem int, count int) error {
	query := `INSERT INTO inventory (user_id, item_id, count) VALUES (?,?,?)`

	_, err := db.Exec(query, idUser, idItem, count)

	if err != nil {
		fmt.Println("ошибка при добавлении предмета в игвентарь предмета", err)
		return err
	}
	fmt.Println("добавлен предмет в инвентарь")
	return err
}

//получение предмета из инвентаря по ид
func getIdInventory(db *sql.DB, id int) (Inventory, error) {
	query := `SELECT * FROM inventory where id=?`

	rows, err := db.Query(query, id)
	if err != nil {
		fmt.Println("fff", err)
		return Inventory{}, err
	}
	var idf Inventory

	for rows.Next() {

		err := rows.Scan(
			&idf.Items,
		)
		if err != nil {
			return Inventory{}, err
		}

	}

	if err := rows.Err(); err != nil {
		return Inventory{}, err
	}
	return idf, nil
}

//получение предмета из инвентаря по типу
func getId_Type_Inventory(db *sql.DB, id int) (Inventory, error) {
	query := `SELECT * FROM inventory where type_id=?`

	rows, err := db.Query(query, id)
	if err != nil {
		fmt.Println("fff", err)
		return Inventory{}, err
	}
	var idf Inventory

	for rows.Next() {

		err := rows.Scan(
			&idf.Items,
		)
		if err != nil {
			return Inventory{}, err
		}

	}

	if err := rows.Err(); err != nil {
		return Inventory{}, err
	}
	return idf, nil
}

//получение всех инвентарей
func getAllInventory(db *sql.DB, user_id int) ([]InventoryDb_228, error) {
	query := `SELECT * FROM inventory where user_id=?`

	rows, err := db.Query(query, user_id)
	if err != nil {
		fmt.Println("fff", err)
		return []InventoryDb_228{}, err
	}
	var inventoris []InventoryDb_228

	for rows.Next() {
		var invetor InventoryDb_228
		err := rows.Scan(
			&invetor.id,
			&invetor.user_id,
			&invetor.item_id,
			&invetor.count,
		)
		if err != nil {
			return []InventoryDb_228{}, err
		}

		inventoris = append(inventoris, invetor)
	}

	if err := rows.Err(); err != nil {
		return []InventoryDb_228{}, err
	}
	return inventoris, nil
}