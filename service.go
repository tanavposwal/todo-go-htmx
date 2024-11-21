package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var DB *sql.DB

func InitDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./todos.db")
	if err != nil {
		log.Fatal(err)
	}
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			status TEXT
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTodo(title string, status string) (int64, error) {
	if DB == nil {
		log.Fatal("Database is not initialized")
	}
	result, err := DB.Exec("INSERT INTO todos (title, status) VALUES (?, ?)", title, status)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func DeleteTodo(id int64) error {
	if DB == nil {
		log.Fatal("Database is not initialized")
	}
	_, err := DB.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}

func ReadTodoList() []Todo {
	if DB == nil {
		log.Fatal("Database is not initialized")
	}
	rows, err := DB.Query("SELECT id, title, status FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	todos := make([]Todo, 0)
	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.Id, &todo.Title, &todo.Status)
		todos = append(todos, todo)
	}
	return todos
}
