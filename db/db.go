package db

import (
	"context"
	"fmt"
	"api-project/modals"
	"github.com/jackc/pgx/v5"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "P@99!db!!73"
	dbname   = "postgres"
)

// Connect establishes a connection to the PostgreSQL database
func Connect() (*pgx.Conn, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbname)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	return conn, nil
}

// CreateAlbumTable creates the albums table if it does not exist
func CreateAlbumTable() error {
	conn, err := Connect()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	// Create a table for albums
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS albums (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255),
		artist VARCHAR(255),
		price DECIMAL
	);`
	_, err = conn.Exec(context.Background(), createTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}

// InsertAlbum inserts a new album into the albums table
func InsertAlbum(title, artist string, price float64) error {
	conn, err := Connect()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	// Insert album into the database
	insertSQL := `INSERT INTO albums (title, artist, price) VALUES ($1, $2, $3)`
	_, err = conn.Exec(context.Background(), insertSQL, title, artist, price)
	if err != nil {
		return fmt.Errorf("failed to insert album: %w", err)
	}

	return nil
}

// GetAlbums retrieves all albums from the database
func GetAlbums() ([]modals.Album, error) {
	conn, err := Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	// Query the database for albums
	rows, err := conn.Query(context.Background(), "SELECT id, title, artist, price FROM albums")
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}
	defer rows.Close()

	var albums []modals.Album
	for rows.Next() {
		var album modals.Album
		err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over rows: %w", err)
	}

	return albums, nil
}