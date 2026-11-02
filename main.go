package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
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
	r := chi.NewRouter()

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		users, err := getUsers(db)
		if err != nil {
			http.Error(w, "ошибка получения пользователей", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			http.Error(w, "ошибка формирования ответа", http.StatusInternalServerError)
		}
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		items, err := getItems(db)
		if err != nil {
			http.Error(w, "ошибка формирования ответа", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(items)
		if err != nil {
			http.Error(w, "ошибка формирования ответа", http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":3000", r)
}
