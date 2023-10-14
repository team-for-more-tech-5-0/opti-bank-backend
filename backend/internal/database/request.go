package database

import (
	"database/sql"
	"github.com/Spawn4real/hackthon_more.tech_5.0.git/internal/models/bank"
)

func GetBanks() ([]bank.Bank, error) {
	// Устанавливаем заголовок ответа на JSON
	db, err := GetDatabase()
	if err != nil {
		panic(err)
	}

	// Запрос к базе данных для получения списка всех банков
	rows, err := db.Query("SELECT * FROM Bank")
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	//Закрытие данных
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var banks []bank.Bank // Предполагается, что у вас есть структура Bank для хранения данных

	for rows.Next() {
		var bank bank.Bank
		if err := rows.Scan(
			&bank.ID,
			&bank.SalePointName,
			&bank.Address,
			&bank.Status,
			&bank.Schedule,
			&bank.Rko,
			&bank.OfficeType,
			&bank.SalePointFormat,
			&bank.Suoavailability,
			&bank.Hasramp,
			&bank.Latitude,
			&bank.Longitude,
			&bank.Metrostation,
			&bank.Distance,
			&bank.Service,
		); err != nil {
			return nil, err
		}
		banks = append(banks, bank)
	}

	return banks, nil
}
