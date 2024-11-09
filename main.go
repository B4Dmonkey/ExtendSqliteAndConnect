package db_connect

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	sqlite "github.com/mattn/go-sqlite3"
)

/*
init registers the sqlite3 driver with the extended functions

	This MUST be run first before any other database operations
*/
func init() {
	log.Println("Extending sqlite3 driver...")

	var CONNECT_ONCE sync.Once

	CONNECT_ONCE.Do(func() {
		sql.Register("sqlite3_extended", &sqlite.SQLiteDriver{
			ConnectHook: func(conn *sqlite.SQLiteConn) error {
				if err := conn.RegisterFunc("uuid", newUUID, false); err != nil {
					return err
				}

				if err := conn.RegisterFunc("current_timestamp", currentTimestamp, false); err != nil {
					return err
				}

				return nil
			},
		})
	})
	log.Println("Extended sqlite3 driver successfully")
}

func currentTimestamp() string { return time.Now().Format("2006-01-02 15:04:05") }

func newUUID() string { return uuid.New().String() }

/*
Connects to the database and returns the connection.

	Errors if the location is empty || the connection fails.
*/
func ConnectToExtendedSqliteDatabase(dbLocation string) (*sql.DB, error) {
	if dbLocation == "" {
		return nil, fmt.Errorf("Database location cannot be empty")
	}

	log.Println("Connecting to database...")

	connection, err := sql.Open("sqlite3_extended", dbLocation)

	if err != nil {
		return nil, fmt.Errorf("Failed to connect the database: %s\n%s", dbLocation, err)
	}

	if err := connection.Ping(); err != nil {
		return nil, fmt.Errorf("Unable to ping database: %s\n%s", dbLocation, err)
	}

	log.Println("Database connected successfully")

	return connection, nil
}
