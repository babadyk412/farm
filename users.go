package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// создание пользователя
func registration(db *sql.DB, login string, password string) {

	query := `INSERT INTO users (login, password,money) VALUES (?,?,?)`

	_, err := db.Exec(query, login, password, 0)
	if err != nil {
		fmt.Println("ошибка при регистрации", err)
		return
	}
	fmt.Println("создан пользователь", login)
}

// вход
func enter(db *sql.DB, login string, pasword string) (User, error) {
	user := User{}

	q := `SELECT *FROM users WHERE login=? AND password=?`
	err := db.QueryRow(q, login, pasword).Scan(&user.Id, &user.Login, &user.Password, &user.Money)
	if err != nil {
		fmt.Println("ошибка", err)
		return User{}, err
	}
	return user, nil
}

// получение пользователя по ид
func getUser(db *sql.DB, id int) (User, error) {
	query := `SELECT * FROM users where id=?`

	rows, err := db.Query(query, id)
	if err != nil {
		fmt.Println("fff", err)
		return User{}, err
	}
	var typey User

	for rows.Next() {

		err := rows.Scan(
			&typey.Id,
			&typey.Login,
			&typey.Password,
			&typey.Money,
		)
		if err != nil {
			return User{}, err
		}

	}

	if err := rows.Err(); err != nil {
		return User{}, err
	}
	return typey, nil
}

// получение всех пользователей
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
