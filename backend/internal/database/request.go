package database

import (
	"database/sql"
	"encoding/json"
	"github.com/team-for-more-tech-5-0/opti-bank-backend.git/internal/models/atm"
	"github.com/team-for-more-tech-5-0/opti-bank-backend.git/internal/models/bank"
	"log"
)

// Получение всех банков из БД
func GetBanks() ([]bank.Bank, error) {
	// Устанавливаем заголовок ответа на JSON
	db, err := GetDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Запрос к базе данных для получения списка всех банков
	rows, err := db.Query("SELECT * FROM bank")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer rows.Close()

	//Результат функции
	var banks []bank.Bank

	//Получаем данные из БД
	for rows.Next() {
		var currentBank bank.Bank
		var openHoursByte []byte
		//sql.NullString - На случай, если переменная в БД пустая
		var rko sql.NullString
		var hasramp sql.NullString
		var metroStation sql.NullString
		var suoavailability sql.NullString
		var kep sql.NullString
		//Т.к эта переменная в БД записана как jsonb, внутри которого еще много jsonb, сконвертируем в тип map[string]map[string]interface{}
		var services json.RawMessage
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
			&currentBank.QueueIndividual,
			&currentBank.QueueBusiness,
			&currentBank.TimeIndividual,
			&currentBank.TimeBusiness,
			&services,
		); err != nil {
			panic(err)
		}
		//decode service
		err = json.Unmarshal(services, &currentBank.Service)
		if err != nil {
			log.Fatal(err)
		}
		//decode OpenHours
		if err := json.Unmarshal(openHoursByte, &currentBank.OpenHours); err != nil {
			panic(err)
		}

		banks = append(banks, currentBank)
	}
	return banks, nil
}

// Получение всех ATM из БД
func GetAtms() ([]atm.Atm, error) {
	db, err := GetDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Запрос к базе данных для получения списка всех банков
	rows, err := db.Query("SELECT * FROM atm")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	//Резлультат функции
	var atms []atm.Atm

	for rows.Next() {
		var currentAtm atm.Atm
		var services json.RawMessage
		if err := rows.Scan(
			&currentAtm.ID,
			&currentAtm.Address,
			&currentAtm.Latitude,
			&currentAtm.Longitude,
			&currentAtm.IsAllDay,
			&services,
			&currentAtm.Time,
			&currentAtm.Queue,
		); err != nil {
			panic(err)
		}

		err = json.Unmarshal(services, &currentAtm.Services)
		if err != nil {
			log.Fatal(err)
		}
		atms = append(atms, currentAtm)
	}
	return atms, nil
}
