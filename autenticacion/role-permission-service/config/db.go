package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		name := os.Getenv("DB_NAME")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, name)

		db, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Fatalf("Error al conectar con PostgreSQL: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Error al hacer ping a la base de datos: %v", err)
		}

		log.Println("✅ Conexión establecida con PostgreSQL")
	})

	return db
}
