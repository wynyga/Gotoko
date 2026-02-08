package connection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/wynyga/gotoko/internal/config"
)

func GetDatabase(conf config.Database) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable Timezone=%s",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Pass,
		conf.Name,
		conf.Tz,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err.Error())
	}
	return db
}
