package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func connect_database() {
	// Строка подключения к базе данных
	connectionString := "user=postgres password=4203 host=localhost port=5432 dbname=postgres sslmode=disable"

	// Устанавливаем соединение с базой данных
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	// Проверяем соединение
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Закрываем соединение при завершении приложения
	defer db.Close()
}
