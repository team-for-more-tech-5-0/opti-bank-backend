package database

import (
	"database/sql"
	"encoding/json"
	"github.com/team-for-more-tech-5-0/opti-bank-backend.git/internal/models/bank"
	"log"
)

func GetBanks() ([]bank.Bank, error) {
	// Устанавливаем заголовок ответа на JSON
	db, err := GetDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Запрос к базе данных для получения списка всех банков
	rows, err := db.Query("SELECT * FROM bank")
	if err != nil {
		log.Fatal(err)
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
		var openHoursByte []byte
		var metroStation sql.NullString
		var serviceByte []byte
		if err := rows.Scan(
			&bank.ID,
			&bank.SalePointName,
			&bank.Address,
			&bank.Status,
			&openHoursByte,
			&bank.Rko,
			&bank.OfficeType,
			&bank.SalePointFormat,
			&bank.Suoavailability,
			&bank.Hasramp,
			&bank.Latitude,
			&bank.Longitude,
			&metroStation,
			&bank.Distance,
			&bank.Kep,
			&bank.MyBranch,
			&serviceByte,
		); err != nil {
			panic(err)
		}
		if err := json.Unmarshal(openHoursByte, &bank.OpenHours); err != nil {
			panic(err)
		}
		err = json.Unmarshal(serviceByte, &bank.Service)
		if err != nil {
			panic(err)
		}

		banks = append(banks, bank)
	}
	return banks, nil
}
