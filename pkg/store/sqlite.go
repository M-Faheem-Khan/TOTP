package store

import (
	"database/sql"
	"fmt"
	"log"
	"m-faheem-khan/totp/pkg/totp"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db *sql.DB
}

func NewStore(dbPath string) (*Store, error) {
	dsn := fmt.Sprintf("file:%s?_journal_mode=WAL&_timeout=5000&_fk=true", dbPath)

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

	// Verify connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Store{db: db}, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) Setup() {
	sqlStmt := `
    CREATE TABLE IF NOT EXISTS totp (
        id INTEGER PRIMARY KEY,
        secret TEXT
    );
    `
	_, err := s.db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table 'totp' created successfully")
}

func (s *Store) GetSecret(id int) *totp.Totp {
	query := `SELECT id, secret FROM totp where id=?;`
	var totp totp.Totp
	err := s.db.QueryRow(query, id).Scan(&totp.Id, &totp.Secret)
	if err != nil {
		log.Fatal(err)
	}
	return &totp
}

func (s *Store) PutSecret(secret string) int64 {
	query := `INSERT INTO totp (secret) VALUES (?);`
	result, err := s.db.Exec(query, secret)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Saved Secert in 'totp' successfully")
	return id
}
