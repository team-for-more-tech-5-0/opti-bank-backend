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
		panic(err)
	}
	defer rows.Close()

	var banks []bank.Bank // Предполагается, что у вас есть структура Bank для хранения данных

	for rows.Next() {
		var currentBank bank.Bank
		var openHoursByte []byte
		var rko sql.NullString
		var hasramp sql.NullString
		var metroStation sql.NullString
		var suoavailability sql.NullString
		var kep sql.NullString
		var jsonbCol1 json.RawMessage
		if err := rows.Scan(
			&currentBank.ID,
			&currentBank.SalePointName,
			&currentBank.Address,
			&currentBank.Status,
			&openHoursByte,
			&rko,
			&currentBank.OfficeType,
			&currentBank.SalePointFormat,
			&suoavailability,
			&hasramp,
			&currentBank.Latitude,
			&currentBank.Longitude,
			&metroStation,
			&currentBank.Distance,
			&kep,
			&currentBank.MyBranch,
			&jsonbCol1,
		); err != nil {
			panic(err)
		}
		err = json.Unmarshal(jsonbCol1, &currentBank.Service)
		if err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal(openHoursByte, &currentBank.OpenHours); err != nil {
			panic(err)
		}

		banks = append(banks, currentBank)
	}
	return banks, nil
}
