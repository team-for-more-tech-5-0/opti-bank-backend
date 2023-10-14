package database

import (
	"net/http"
)

func GetBanks(w http.ResponseWriter, r *http.Request) {

	// // Устанавливаем заголовок ответа на JSON
	// w.Header().Set("Content-Type", "application/json")

	// // Открываем соединение с базой данных (предполагается, что у вас уже есть соединение)
	// db, err := sql.Open("postgres", "your_db_connection_string_here")
	// if err != nil {
	//     http.Error(w, err.Error(), http.StatusInternalServerError)
	//     return
	// }
	// defer db.Close()

	// // Запрос к базе данных для получения списка всех банков
	// rows, err := db.Query("SELECT * FROM Bank")
	// if err != nil {
	//     http.Error(w, err.Error(), http.StatusInternalServerError)
	//     return
	// }
	// defer rows.Close()

	// var banks []Bank // Предполагается, что у вас есть структура Bank для хранения данных

	// for rows.Next() {
	//     var bank Bank
	//     if err := rows.Scan(&bank.bank_id, &bank.salePointName, &bank.address, /* ... и так далее */); err != nil {
	//         http.Error(w, err.Error(), http.StatusInternalServerError)
	//         return
	//     }
	//     banks = append(banks, bank)
	// }

	// // Преобразуем список банков в формат JSON и отправляем в ответе
	// if err := json.NewEncoder(w).Encode(banks); err != nil {
	//     http.Error(w, err.Error(), http.StatusInternalServerError)
	//     return
	// }
}
