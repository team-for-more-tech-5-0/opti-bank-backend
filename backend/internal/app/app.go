package app

import (
	"github.com/Spawn4real/hackthon_more.tech_5.0.git/internal/database"
	"github.com/Spawn4real/hackthon_more.tech_5.0.git/internal/transport"
)

func Start() {
	transport.Transport()
	database.CloseConnection()
}
