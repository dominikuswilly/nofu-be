package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"nofu-be/config"
)

var DB *sql.DB

func ConnectDB(cfg *config.Config) {
	var err error
	DB, err = sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Gagal koneksi ke PostgreSQL:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Gagal ping ke database:", err)
	}

	log.Println("Berhasil terhubung ke PostgreSQL")
}
