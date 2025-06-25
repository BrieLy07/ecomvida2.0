package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	db   *sql.DB
	once sync.Once
)

func ConectarDB() *sql.DB {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error cargando .env")
		}

		usuario := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		nombre := os.Getenv("DB_NAME")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", usuario, password, host, port, nombre)

		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Error al conectar con la base de datos: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Error al hacer ping a la base de datos: %v", err)
		}

		log.Println("✅ Conexión establecida con MySQL")
	})

	return db
}
