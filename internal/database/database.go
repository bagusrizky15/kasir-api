package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Untuk membatasi koneksi database agar db ga kerepotan
	db.SetMaxOpenConns(25)

	// Untuk menyimpan 5 koneksi nganggur supaya bisa dipakai ulang tanpa membuka koneksi baru
	db.SetMaxIdleConns(5)

	log.Print("Database connected success")
	return db, nil

}
