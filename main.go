package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id       int
	Login    string
	Password string
	Money    int
}

type Item struct {
	Id      int
	Name    string
	Type_id int
	Price   int
}

type Type_id struct {
	Id   int
	Name string
}

type Inventory struct {
	Name  string
	Items []Item
}

type InventoryDb_228 struct {
	id      int
	user_id int
	item_id int
	count   int
}

func main() {

	db, err := sql.Open("sqlite3", "farmTables.db")
	if err != nil {
		fmt.Println("ошибка подключения:", err)
		return
	}
	defer db.Close()

	fmt.Println("что делать")

	var vibor int
	fmt.Scan(&vibor)

	if vibor == 1 {
		registration(db)
	}
	if vibor == 2 {
		user, err := enter(db)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(user)
	}
	if vibor == 3 {
		users, err := getUsers(db)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, i := range users {
			fmt.Println(i.Login)
		}
	}
	if vibor == 22 {
		addItem(db)
	}
	if vibor == 222 {
		items, err := getItems(db)
		if err != nil {
			fmt.Println("ошибка при получении предметов", err)
			return
		}
		for _, item := range items {
			fmt.Printf("название: %s; цена: %d", item.Name, item.Price)
		}
	}
	if vibor == 4 {
		addItemInventory(db)
	}
	if vibor == 5 {
		addItemType(db)
	}
	if vibor == 6 {
		addItem(db)
	}
	if vibor == 7 {
		fmt.Println("ид юзера")
		var user_id int
		fmt.Scan(&user_id)
		inventory, err := getAllInventory(db, user_id)
		if err != nil {
			fmt.Println("ошибка при получении инвентарей", err)
			return
		}
		for _, inv := range inventory {
			item, err := getIdItem(db, inv.item_id)
			fmt.Println(inv.item_id)
			fmt.Println(item.Name)
			if err != nil {
				fmt.Println("ошибка при получении предмета по ид", err)
				return
			}
			fmt.Printf("ид %d,%s,%d", user_id, item.Name, inv.count)
		}

	}

}

func registration(db *sql.DB) {
	fmt.Println("введи логин")
	var login string
	fmt.Scan(&login)

	fmt.Println("введи пароль")
	var pasword string
	fmt.Scan(&pasword)

	user := User{
		Login:    login,
		Password: pasword,
	}

	query := `INSERT INTO users (login, password,money) VALUES (?,?,?)`

	fmt.Println("создан пользователь", user.Login)

	_, err := db.Exec(query, login, pasword, 0)
	if err != nil {
		fmt.Println("ошибка при регистрации", err)
		return
	}
}

func getUsers(db *sql.DB) ([]User, error) {
	query := `SELECT * FROM users`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("fff", err)
		return []User{}, err
	}
	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(
			&user.Id,
			&user.Login,
			&user.Password,
			&user.Money,
		)
		if err != nil {
			return []User{}, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return []User{}, err
	}
	return users, nil
}

func enter(db *sql.DB) (User, error) {
	fmt.Println("логин для входа")
	var login string
	fmt.Scan(&login)

	fmt.Println("пароль для входа")
	var pasword string
	fmt.Scan(&pasword)

	user := User{}

	q := `SELECT *FROM users WHERE login=? AND password=?`
	err := db.QueryRow(q, login, pasword).Scan(&user.Id, &user.Login, &user.Password, &user.Money)
	if err != nil {
		fmt.Println("ошибка", err)
		return User{}, err
	}
	return user, nil
}

func addItem(db *sql.DB) {
	fmt.Println("название")
	var name string
	fmt.Scan(&name)

	fmt.Println("цена")
	var price int
	fmt.Scan(&price)

	fmt.Println("Typeid предмета")
	var typeId int
	fmt.Scan(&typeId)

	item := Item{
		Name:    name,
		Price:   price,
		Type_id: typeId,
	}

	query := `INSERT INTO items (name, price, type_id) VALUES (?,?,?)`

	fmt.Println("создан предмет", item.Name)

	_, err := db.Exec(query, name, price, typeId)
	if err != nil {
		fmt.Println("ошибка при создание предмета", err)
		return
	}
}

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

func addItemType(db *sql.DB) {
	fmt.Println("название")
	var name string
	fmt.Scan(&name)

	typey := Type_id{
		Name: name,
	}

	query := `INSERT INTO items_type (name) VALUES (?)`

	fmt.Println("создан тип", typey.Name)

	_, err := db.Exec(query, name)
	if err != nil {
		fmt.Println("ошибка при создание типа", err)
		return
	}
}

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

// ////////////////////////////
func getTypeId(db *sql.DB) (Type_id, error) {
	query := `SELECT * FROM items where id=?`

	fmt.Println("по какому айди найти")
	var id int
	fmt.Scan(&id)

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

func addItemInventory(db *sql.DB) error {
	fmt.Println("ид пользователя")
	var idUser int
	fmt.Scan(&idUser)

	fmt.Println("ид предмета")
	var idItem int
	fmt.Scan(&idItem)

	fmt.Println("кол во")
	var count int
	fmt.Scan(&count)

	query := `INSERT INTO inventory (user_id, item_id, count) VALUES (?,?,?)`

	_, err := db.Exec(query, idUser, idItem, count)

	if err != nil {
		fmt.Println("ошибка при добавлении предмета в игвентарь предмета", err)
		return err
	}
	fmt.Println("добавлен предмет в инвентарь")
	return err
}

func getIdInventory(db *sql.DB) (Inventory, error) {
	query := `SELECT * FROM inventory where id=?`

	fmt.Println("по какому айди найти")
	var id int
	fmt.Scan(&id)

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

func getId_User_Inventory(db *sql.DB) (Inventory, error) {
	query := `SELECT * FROM inventory where user_id=?`

	fmt.Println("по какому айди пользователя найти")
	var id int
	fmt.Scan(&id)

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

func getId_Type_Inventory(db *sql.DB) (Inventory, error) {
	query := `SELECT * FROM inventory where type_id=?`

	fmt.Println("по какому типу предмета найти")
	var id string
	fmt.Scan(&id)

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
