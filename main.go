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
		fmt.Println("введи логин")
		var login string
		fmt.Scan(&login)

		fmt.Println("введи пароль")
		var pasword string
		fmt.Scan(&pasword)

		registration(db, login, pasword)
	}
	if vibor == 2 {
		fmt.Println("логин для входа")
		var login string
		fmt.Scan(&login)

		fmt.Println("пароль для входа")
		var pasword string
		fmt.Scan(&pasword)

		user, err := enter(db, login, pasword)
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
		fmt.Println("название")
		var name string
		fmt.Scan(&name)

		fmt.Println("цена")
		var price int
		fmt.Scan(&price)

		fmt.Println("Typeid предмета")
		var typeId int
		fmt.Scan(&typeId)

		addItem(db, name, price, typeId)
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
		fmt.Println("ид пользователя")
		var idUser int
		fmt.Scan(&idUser)

		fmt.Println("ид предмета")
		var idItem int
		fmt.Scan(&idItem)

		fmt.Println("кол во")
		var count int
		fmt.Scan(&count)

		addItemInventory(db, idUser, idItem, count)
	}
	if vibor == 5 {
		fmt.Println("название")
		var name string
		fmt.Scan(&name)

		addItemType(db, name)
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
		user, err := getUser(db, user_id)
		if err != nil {
			fmt.Println("ошибка при получении 1пользователя", err)
			return
		}
		for _, inv := range inventory {
			item, err := getIdItem(db, inv.item_id)
			if err != nil {
				fmt.Println("ошибка при получении предмета по ид", err)
				return
			}

			fmt.Printf("имя: %s, предмет: %s, количество: %d\n", user.Login, item.Name, inv.count)
		}

	}

}
