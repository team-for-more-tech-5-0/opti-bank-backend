package database

func CloseConnection() {
	// Закрываем соединение при завершении приложения
	db, err := GetDatabase()
	if err != nil {
		panic(err)
	}

	err = db.Close()
	if err != nil {
		panic(err)
	}
}
