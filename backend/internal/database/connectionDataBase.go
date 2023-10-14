package database

import (
	"database/sql"
	"github.com/Spawn4real/hackthon_more.tech_5.0.git/internal/config"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	once sync.Once
	db   *sql.DB
	err  error
)

func GetDatabase() (*sql.DB, error) {
	once.Do(func() {
		// Строка подключения к базе данных
		connectionString := config.ConnectionDataBaseString

		// Устанавливаем соединение с базой данных
		db, err := sql.Open(config.DataBaseName, connectionString)
		if err != nil {
			log.Fatal(err)
		}

		// Проверяем соединение
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}
	})

	return db, err
}
